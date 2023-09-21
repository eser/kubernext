package shared

import (
	"github.com/eser/go-service/pkg/infra/httpserv"
)

func RequestIdMiddleware() httpserv.HandlerFunc {
	return func(ctx *httpserv.Context) {
		requestId := GenerateUniqueId()

		ctx.Set("requestId", requestId)
		ctx.Header("Request-Id", requestId)

		ctx.Next()
	}
}
