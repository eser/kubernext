package pkg

import (
	"github.com/eser/go-service/pkg/entities"
	"github.com/eser/go-service/pkg/healthcheck"
	"github.com/eser/go-service/pkg/infra"
	"github.com/eser/go-service/pkg/infra/httpserv"
	"github.com/eser/go-service/pkg/shared"
	"go.uber.org/fx"
)

func NewRootRouter(hs *httpserv.HttpServ) *httpserv.RouterGroup {
	routes := hs.Engine.Group("/")
	routes.Use(shared.ResponseTimeMiddleware()) // FIXME(@eser): must be the first middleware
	routes.Use(shared.RequestIdMiddleware())
	routes.Use(shared.ErrorHandlerMiddleware())

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

func Run() {
	modules := fx.Options(
		fx.WithLogger(infra.GetFxLogger),
		infra.InfraModule,
		AppModule,
	)

	fx.New(modules).Run()
}
