package infra

import (
	"github.com/eser/go-service/pkg/infra/config"
	"github.com/eser/go-service/pkg/infra/httpserv"
	"github.com/eser/go-service/pkg/infra/log"
	"github.com/eser/go-service/pkg/infra/mongo"
	"github.com/eser/go-service/pkg/infra/redis"
	"go.uber.org/fx"
)

var (
	GetFxLogger = log.GetFxLogger
	Module      = fx.Module(
		"infra",
		fx.Provide(
			config.NewConfig,
			log.NewRuntimeContext,
			log.NewLogger,
			httpserv.NewHttpServ,
			redis.NewRedisClient,
			mongo.NewMongoClient,
		),
		fx.Invoke(
			httpserv.RegisterHttpLogger,
			httpserv.RegisterHooks,
		),
	)
)
