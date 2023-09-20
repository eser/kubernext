package httpserv

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eser/go-service/lib/config"
	"github.com/eser/go-service/lib/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type (
	Engine      = gin.Engine
	Context     = gin.Context
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

	// cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowCredentials = true

	engine.Use(cors.New(corsConfig))

	srv := &http.Server{
		Addr:    ":" + conf.Port,
		Handler: engine,
	}

	return &HttpServ{srv, engine}, nil
}

func RegisterHooks(lc fx.Lifecycle, hs *HttpServ, conf *config.Config, logger *log.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("HttpServ starting", log.String("env", conf.Env), log.String("port", conf.Port))

				if err := hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Fatal("HttpServ ListenAndServe error", log.Error(err))

					// return
				}

				// logger.Info("HttpServ.ListenAndServe() is stopped")
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			// Wait for interrupt signal to gracefully shutdown the server with
			// a timeout of 5 seconds.
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

			logger.Info("HttpServ shutting down...")
			<-quit

			// The context is used to inform the server it has 5 seconds to finish
			// the request it is currently handling
			ctx, cancel := context.WithTimeout(context.Background(), GracefulShutdownTimeout)
			defer cancel()

			if err := hs.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
				logger.Fatal("HttpServ forced to shutdown: ", log.Error(err))

				return err
			}

			logger.Info("HttpServ exiting")

			return nil
		},
	})
}
