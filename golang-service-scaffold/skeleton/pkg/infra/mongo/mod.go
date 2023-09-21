package mongo

import (
	"context"
	"time"

	"github.com/eser/go-service/pkg/infra/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type (
	MongoClient = mongo.Client
)

func NewMongoClient(conf *config.Config) (*MongoClient, error) {
	if conf.MongoAddr == "" {
		return nil, nil
	}

	mongoOptions := options.Client().ApplyURI(conf.MongoAddr)
	mongoOptions.WriteConcern = writeconcern.Majority()
	if conf.MongoMaxPoolSize > 0 {
		mongoOptions.SetMaxPoolSize(conf.MongoMaxPoolSize)
	}

	var ctx context.Context
	if conf.MongoConnTimeout > 0 {
		var ctxCancel context.CancelFunc
		ctx, ctxCancel = context.WithTimeout(context.Background(), time.Duration(conf.MongoConnTimeout)*time.Second)
		defer ctxCancel()
	} else {
		ctx = context.Background()
	}

	mongoClient, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		return nil, err
	}

	if conf.MongoConnCheck {
		err = mongoClient.Ping(context.Background(), nil)
		if err != nil {
			return nil, err
		}
	}

	return mongoClient, nil
}
