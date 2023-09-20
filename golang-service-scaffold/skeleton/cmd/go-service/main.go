package main

import (
	"github.com/eser/go-service/lib"
	"github.com/eser/go-service/pkg"
	"go.uber.org/fx"
)

func main() {
	modules := fx.Options(
		fx.WithLogger(lib.GetFxLogger),
		lib.InfraModule,
		pkg.AppModule,
	)
	fx.New(modules).Run()
}
