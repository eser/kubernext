package infra

import (
	"github.com/eser/go-service/lib/config"
	"github.com/eser/go-service/lib/httpserv"
	"github.com/eser/go-service/lib/log"
	"go.uber.org/fx"
)

var (
	GetFxLogger = log.GetFxLogger
	Module      = fx.Module(
		"infra",
		fx.Provide(
			log.NewLogger,
			config.NewConfig,
			httpserv.NewHttpServ,
		),
		fx.Invoke(
			httpserv.RegisterHttpLogger,
			httpserv.RegisterHooks,
		),
	)
)
