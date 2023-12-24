// This file provides log handling for httpserv.
// Code based on ginrus and ginzap packages.
package httpserv

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/eser/go-service/pkg/infra/log"
)

type (
	Fn        func(c *Context) []log.Attr
	LogConfig struct {
		TimeFormat string
		UTC        bool
		SkipPaths  []string
		Context    Fn
	}
	CustomWriter struct {
		Logger *log.Logger
	}
)

func generateLogFields(
	ctx *Context, conf *LogConfig, path string, query string, latency time.Duration, end time.Time,
) []any {
	fields := []any{
		log.String("event", "http-request"),
		log.Int("status", ctx.Writer.Status()),
		log.String("method", ctx.Request.Method),
		log.String("path", path),
		log.String("query", query),
		log.String("ip", ctx.ClientIP()),
		log.String("user-agent", ctx.Request.UserAgent()),
		log.Duration("latency", latency),
	}

	if conf.TimeFormat != "" {
		fields = append(fields, log.String("time", end.Format(conf.TimeFormat)))
	}

	if conf.Context != nil {
		for _, attr := range conf.Context(ctx) {
			fields = append(fields, attr)
		}
	}

	return fields
}

func LoggerMiddleware(logger *log.Logger, timeFormat string, utc bool) HandlerFunc {
	conf := &LogConfig{TimeFormat: timeFormat, UTC: utc}
	skipPaths := make(map[string]bool, len(conf.SkipPaths))

	for _, path := range conf.SkipPaths {
		skipPaths[path] = true
	}

	return func(ctx *Context) {
		start := time.Now()

		// some evil middlewares modify this values
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery

		ctx.Next()

		if _, ok := skipPaths[path]; !ok {
			end := time.Now()
			latency := end.Sub(start)

			if conf.UTC {
				end = end.UTC()
			}

			fields := generateLogFields(ctx, conf, path, query, latency, end)

			if len(ctx.Errors) > 0 {
				// Append error field if this is an erroneous request.
				for _, e := range ctx.Errors.Errors() {
					logger.Error(e, fields...)
				}

				return
			}

			logger.Info(path, fields...)
		}
	}
}

func defaultHandleRecovery(ctx *Context, _err any) {
	ctx.AbortWithStatus(http.StatusInternalServerError)
}

// Check for a broken connection, as it is not really a
// condition that warrants a panic stack trace.
func checkBrokenPipe(err any) bool {
	if ne, ok := err.(*net.OpError); ok {
		var se *os.SyscallError

		if errors.As(ne, &se) {
			errLower := strings.ToLower(se.Error())

			if strings.Contains(errLower, "broken pipe") || strings.Contains(errLower, "connection reset by peer") {
				return true
			}
		}
	}

	return false
}

func RecoveryWithLoggerMiddleware(logger *log.Logger, stack bool) HandlerFunc {
	recovery := defaultHandleRecovery

	return func(ctx *Context) {
		defer func() {
			err := recover()

			if err == nil {
				return
			}

			httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
			brokenPipe := checkBrokenPipe(err)

			if brokenPipe {
				logger.Error(ctx.Request.URL.Path,
					log.Any("error", err),
					log.String("request", string(httpRequest)),
				)

				// If the connection is dead, we can't write a status to it.
				errAsError, ok := err.(error)
				if ok {
					ctx.Error(errAsError) //nolint:errcheck
				}

				ctx.Abort()

				return
			}

			if stack {
				logger.Error("[Recovery from panic]",
					log.Time("time", time.Now()),
					log.Any("error", err),
					log.String("request", string(httpRequest)),
					log.String("stack", string(debug.Stack())),
				)
			} else {
				logger.Error("[Recovery from panic]",
					log.Time("time", time.Now()),
					log.Any("error", err),
					log.String("request", string(httpRequest)),
				)
			}

			recovery(ctx, err)
		}()

		ctx.Next()
	}
}

func RegisterHttpLogger(logger *log.Logger, hs *HttpServ) {
	hs.Engine.Use(
		//   - Logs all requests, like a combined access and error log.
		//   - Logs to stdout.
		//   - RFC3339 with UTC time format.
		LoggerMiddleware(logger, time.RFC3339, true),

		// Logs all panic to error log
		//   - stack means whether output the stack info.
		RecoveryWithLoggerMiddleware(logger, true),
	)
}

func (w *CustomWriter) Write(p []byte) (int, error) {
	str := string(p)

	if strings.HasPrefix(str, "[GIN-debug]") {
		w.Logger.Debug(str)
	} else {
		w.Logger.Info(str)
	}

	return len(str), nil
}
