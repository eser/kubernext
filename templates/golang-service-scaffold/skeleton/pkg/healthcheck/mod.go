package healthcheck

import (
	"github.com/eser/go-service/pkg/infra/httpserv"
	"github.com/tavsec/gin-healthcheck/checks"
	"github.com/tavsec/gin-healthcheck/config"
	"github.com/tavsec/gin-healthcheck/controllers"
	"go.uber.org/fx"
)

var Module = fx.Module( //nolint:gochecknoglobals
	"healthcheck",
	fx.Invoke(
		RegisterRoutes,
	),
)

func RegisterRoutes(router *httpserv.RouterGroup) {
	conf := config.DefaultConfig()
	conf.HealthPath = "/health-check"

	checks := []checks.Check{}

	router.GET(conf.HealthPath, controllers.HealthcheckController(checks, conf))
}
