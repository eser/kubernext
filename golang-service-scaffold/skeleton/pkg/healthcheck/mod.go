package healthcheck

import (
	"github.com/eser/go-service/pkg/infra/httpserv"
	"github.com/tavsec/gin-healthcheck/checks"
	"github.com/tavsec/gin-healthcheck/config"
	"github.com/tavsec/gin-healthcheck/controllers"
)

func RegisterRoutes(rg *httpserv.RouterGroup) {
	conf := config.Config{HealthPath: "/health-check"}

	checks := []checks.Check{}

	rg.GET(conf.HealthPath, controllers.HealthcheckController(checks, conf))
}
