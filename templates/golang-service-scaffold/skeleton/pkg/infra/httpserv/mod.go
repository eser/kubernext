package httpserv

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/eser/go-service/pkg/infra/config"
	"github.com/eser/go-service/pkg/infra/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type (
	Engine      = gin.Engine
	Context     = gin.Context
	Request     = http.Request
	Response    = http.Response
	HandlerFunc = gin.HandlerFunc
	RouterGroup = gin.RouterGroup
	H           = gin.H
)

type HttpServ struct {
	*http.Server

	Engine *Engine
}

const (
	GracefulShutdownTimeout = 5 * time.Second
)

func NewHttpServ(conf *config.Config, logger *log.Logger) (*HttpServ, error) {
	gin.DefaultErrorWriter = &CustomWriter{logger}
	gin.DefaultWriter = &CustomWriter{logger}

	// instantiation
	if conf.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	// engine.Use(gin.Logger())
	engine.Use(ResponseWriterHandler)

	// cors
	corsConfig := cors.DefaultConfig()

	if conf.CorsOrigin == "" {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowOrigins = []string{conf.CorsOrigin}
	}

	if conf.CorsStrictHeaders {
		corsConfig.AllowHeaders = []string{
			"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization",
		}
	} else {
		corsConfig.AllowHeaders = []string{"*"}
	}

	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true

	engine.Use(cors.New(corsConfig))

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	srv := &http.Server{
		ReadHeaderTimeout: 5 * time.Second,  //nolint:gomnd
		ReadTimeout:       10 * time.Second, //nolint:gomnd
		WriteTimeout:      10 * time.Second, //nolint:gomnd
		Addr:              ":" + conf.Port,
		Handler:           engine,
	}

	return &HttpServ{srv, engine}, nil
}

func RegisterHooks(lc fx.Lifecycle, hs *HttpServ, conf *config.Config, logger *log.Logger) { //nolint:varnamelen
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("HttpServ is starting...", log.String("env", conf.Env), log.String("port", conf.Port))

			ln, lnErr := net.Listen("tcp", hs.Addr) //nolint:varnamelen
			if lnErr != nil {
				return fmt.Errorf("HttpServ Net Listen error: %w", lnErr)
			}

			go func() {
				if err := hs.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error("HttpServ Serve error", log.ErrorObject(err))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("HttpServ is stopping...")

			shutdownCtx, cancel := context.WithTimeout(ctx, GracefulShutdownTimeout)
			defer cancel()

			err := hs.Shutdown(shutdownCtx)
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				return fmt.Errorf("HttpServ forced to shutdown: %w", err)
			}

			<-shutdownCtx.Done()
			logger.Info("HttpServ has stopped...")

			return nil
		},
	})
}
