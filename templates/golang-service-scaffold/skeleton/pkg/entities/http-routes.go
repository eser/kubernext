package entities

import (
	"net/http"
	"time"

	"github.com/eser/go-service/pkg/infra/httpserv"
	"github.com/eser/go-service/pkg/infra/log"
	"github.com/eser/go-service/pkg/shared"
)

// definition

type EntitiesHttpRoutes struct {
	Logger          *log.Logger
	EntitiesService *EntitiesService
}

func NewEntitiesHttpRoutes(logger *log.Logger, entitiesService *EntitiesService) *EntitiesHttpRoutes {
	return &EntitiesHttpRoutes{
		Logger:          logger,
		EntitiesService: entitiesService,
	}
}

// get

type GetActionRequest struct {
	Id string `json:"id"`
}

type GetActionResponse struct {
	Record *Model `json:"record"`
}

func (s *EntitiesHttpRoutes) GetAction(ctx *httpserv.Context) {
	request := new(GetActionRequest)

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	audit := shared.AuditRecord{
		RequestedBy: ctx.GetString("uid"),
		RequestedAt: time.Now(),
	}
	record, err := s.EntitiesService.GetRecord(request.Id, &audit)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := GetActionResponse{
		Record: record,
	}

	ctx.JSON(http.StatusOK, response)
}
