// This file provides log handling for httpserv.
// Code based on ginrus and ginzap packages.
package httpserv

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/eser/go-service/pkg/infra/log"
	"github.com/gin-gonic/gin"
)

type (
	Fn        func(c *gin.Context) []log.Field
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

func LoggerMiddleware(logger *log.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	conf := &LogConfig{TimeFormat: timeFormat, UTC: utc}
	skipPaths := make(map[string]bool, len(conf.SkipPaths))

	for _, path := range conf.SkipPaths {
		skipPaths[path] = true
	}

	return func(c *gin.Context) {
		start := time.Now()

		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		if _, ok := skipPaths[path]; !ok {
			end := time.Now()
			latency := end.Sub(start)

			if conf.UTC {
				end = end.UTC()
			}

			fields := []log.Field{
				log.String("event", "http-request"),
				log.Int("status", c.Writer.Status()),
				log.String("method", c.Request.Method),
				log.String("path", path),
				log.String("query", query),
				log.String("ip", c.ClientIP()),
				log.String("user-agent", c.Request.UserAgent()),
				log.Duration("latency", latency),
			}

			if conf.TimeFormat != "" {
				fields = append(fields, log.String("time", end.Format(conf.TimeFormat)))
			}

			if conf.Context != nil {
				fields = append(fields, conf.Context(c)...)
			}

			if len(c.Errors) > 0 {
				// Append error field if this is an erroneous request.
				for _, e := range c.Errors.Errors() {
					logger.Error(e, fields...)
				}
			} else {
				logger.Info(path, fields...)
			}
		}
	}
}

func defaultHandleRecovery(c *gin.Context, err interface{}) {
	c.AbortWithStatus(http.StatusInternalServerError)
}

func RecoveryWithLoggerMiddleware(logger *log.Logger, stack bool) gin.HandlerFunc {
	recovery := defaultHandleRecovery

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						log.Any("error", err),
						log.String("request", string(httpRequest)),
					)

					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()

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

				recovery(c, err)
			}
		}()

		c.Next()
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

func (w *CustomWriter) Write(p []byte) (n int, err error) {
	str := string(p)

	if strings.HasPrefix(str, "[GIN-debug]") {
		w.Logger.Debug(str)
	} else {
		w.Logger.Info(str)
	}

	return len(str), nil
}
