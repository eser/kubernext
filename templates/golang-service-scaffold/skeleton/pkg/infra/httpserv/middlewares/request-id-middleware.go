package middlewares

import (
	"github.com/eser/go-service/pkg/infra/httpserv"
	"github.com/eser/go-service/pkg/infra/lib"
)

func RequestIdMiddleware() httpserv.HandlerFunc {
	return func(ctx *httpserv.Context) {
		requestId := lib.GenerateUniqueId()

		ctx.Set("requestId", requestId)
		ctx.Header("X-Request-Id", requestId)

		ctx.Next()
	}
}
