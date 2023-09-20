package entities

import (
	"net/http"
	"time"

	"github.com/eser/go-service/lib/httpserv"
	"github.com/eser/go-service/pkg/shared"
)

type CreateActionRequest struct {
	Fullname string `json:"fullname" binding:"required"`
}

func CreateAction(c *httpserv.Context) {
	request := new(CreateActionRequest)

	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	service := NewService()
	audit := shared.AuditRecord{
		RequestedBy: c.GetString("uid"),
		RequestedAt: time.Now(),
	}
	err := service.CreateRecord(CreateRecordDto{Fullname: request.Fullname}, &audit)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}
