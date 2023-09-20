package healthCheck

import (
	"github.com/eser/go-service/lib/httpserv"
	"github.com/tavsec/gin-healthcheck/checks"
	"github.com/tavsec/gin-healthcheck/config"
	"github.com/tavsec/gin-healthcheck/controllers"
)

func RegisterRoutes(rg *httpserv.RouterGroup) {
	routes := rg // .Group("/health-check")

	conf := config.Config{HealthPath: "/health-check"}

	checks := []checks.Check{}

	routes.GET(conf.HealthPath, controllers.HealthcheckController(checks, conf))
}
