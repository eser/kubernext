package shared

import (
	"strconv"
	"time"

	"github.com/eser/go-service/pkg/infra/httpserv"
)

func ResponseTimeMiddleware() httpserv.HandlerFunc {
	return func(ctx *httpserv.Context) {
		startTime := time.Now()

		ctx.Next()

		duration := int(time.Since(startTime).Milliseconds())

		ctx.Header("Request-Id", strconv.Itoa(duration))
	}
}
