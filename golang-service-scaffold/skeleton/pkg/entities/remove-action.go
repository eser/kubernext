package entities

import (
	"net/http"
	"time"

	"github.com/eser/go-service/lib/httpserv"
	"github.com/eser/go-service/pkg/shared"
)

type RemoveActionRequest struct {
	Id string `json:"id" binding:"required"`
}

func RemoveAction(c *httpserv.Context) {
	request := new(RemoveActionRequest)

	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	service := NewService()
	audit := shared.AuditRecord{
		RequestedBy: c.GetString("uid"),
		RequestedAt: time.Now(),
	}
	err := service.RemoveRecord(RemoveRecordDto{Id: request.Id}, &audit)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
