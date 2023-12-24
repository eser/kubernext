package entities

import (
	"github.com/eser/go-service/pkg/infra/config"
	"github.com/eser/go-service/pkg/infra/httpserv"
	"github.com/eser/go-service/pkg/infra/httpserv/middlewares"
	"go.uber.org/fx"
)

var Module = fx.Module( //nolint:gochecknoglobals
	"entities",
	fx.Provide(
		NewEntitiesService,
		NewEntitiesHttpRoutes,
	),
	fx.Invoke(
		RegisterRoutes,
	),
)

func RegisterRoutes(
	conf *config.Config,
	router *httpserv.RouterGroup,
	entitiesHttpRoutes *EntitiesHttpRoutes,
) {
	isDevelopmentEnv := false
	if conf.Env != "production" {
		isDevelopmentEnv = true
	}

	routes := router.Group("/entities")
	routes.Use(middlewares.AuthMiddleware(conf.JwtSignature, false, isDevelopmentEnv))

	routes.POST("/get", entitiesHttpRoutes.GetAction)
	// routes.POST("/create", CreateAction)
	// routes.POST("/update", UpdateAction)
	// routes.POST("/remove", RemoveAction)
}
