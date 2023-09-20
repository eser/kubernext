package entities

import (
	"github.com/eser/go-service/lib/httpserv"
	"github.com/eser/go-service/pkg/shared"
)

func RegisterRoutes(rg *httpserv.RouterGroup) {
	routes := rg.Group("/entities")
	routes.Use(shared.AuthMiddleware())

	routes.POST("/get", GetAction)
	routes.POST("/create", CreateAction)
	routes.POST("/update", UpdateAction)
	routes.POST("/remove", RemoveAction)
}
