package pkg

import (
	"github.com/eser/go-service/lib/httpserv"
	"github.com/eser/go-service/pkg/entities"
	"github.com/eser/go-service/pkg/healthcheck"
	"github.com/eser/go-service/pkg/shared"
	"go.uber.org/fx"
)

func NewRootRouter(hs *httpserv.HttpServ) *httpserv.RouterGroup {
	routes := hs.Engine.Group("/")
	routes.Use(shared.ErrorHandler())

	return routes
}

var (
	AppModule = fx.Module(
		"app",
		fx.Provide(
			NewRootRouter,
		),
		fx.Invoke(
			entities.RegisterRoutes,
			healthcheck.RegisterRoutes,
		),
	)
)
