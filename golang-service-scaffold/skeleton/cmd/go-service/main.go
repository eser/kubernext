package main

import (
	// "fmt"

	infra "github.com/eser/go-service/lib"
	// "github.com/eser/go-service/pkg/shared"
	"go.uber.org/fx"
)

func main() {
	// routes := server.Group("/")
	// routes.Use(middlewares.ErrorHandler())
	// httplogs.Bind(routes, logger)

	// modules.RegisterRoutes(routes)

	modules := fx.Options(
		fx.WithLogger(infra.GetFxLogger),
		infra.Module,
	)
	fx.New(modules).Run()
}
