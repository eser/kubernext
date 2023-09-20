package modules

import (
	"github.com/eser/go-service/lib/httpserv"
	"github.com/eser/go-service/pkg/entities"
	healthCheck "github.com/eser/go-service/pkg/health-check"
)

func RegisterRoutes(rg *httpserv.RouterGroup) {
	entities.RegisterRoutes(rg)
	healthCheck.RegisterRoutes(rg)
}
