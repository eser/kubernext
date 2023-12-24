package pkg

import (
	"github.com/eser/go-service/pkg/entities"
	"github.com/eser/go-service/pkg/healthcheck"
	"github.com/eser/go-service/pkg/infra"
	"github.com/eser/go-service/pkg/infra/httpserv"
	"github.com/eser/go-service/pkg/infra/httpserv/middlewares"
	"go.uber.org/fx"
)

var Module = fx.Module( //nolint:gochecknoglobals
	"app",
	fx.Provide(
		NewAppRootHttpRouter,
	),
	healthcheck.Module,
	entities.Module,
)

func NewAppRootHttpRouter(hs *httpserv.HttpServ) *httpserv.RouterGroup {
	routes := hs.Engine.Group("/")
	routes.Use(middlewares.ErrorHandlerMiddleware())
	routes.Use(middlewares.ResponseTimeMiddleware())
	routes.Use(middlewares.RequestIdMiddleware())
	routes.Use(middlewares.IpResolverMiddleware())

	return routes
}

func RunService() {
	fx.New(
		fx.WithLogger(infra.GetFxLogger),
		infra.Module,
		Module,
	).Run()
}
