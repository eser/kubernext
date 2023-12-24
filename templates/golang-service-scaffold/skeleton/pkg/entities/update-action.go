package entities

// import (
// 	"net/http"
// 	"time"

// 	"github.com/eser/go-service/pkg/infra/httpserv"
// 	"github.com/eser/go-service/pkg/shared"
// )

// type UpdateActionRequest struct {
// 	Id       string `json:"id" binding:"required"`
// 	Fullname string `json:"fullname" binding:"required"`
// }

// func UpdateAction(c *httpserv.Context) {
// 	request := new(UpdateActionRequest)

// 	if err := c.ShouldBind(&request); err != nil {
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	service := NewService()
// 	audit := shared.AuditRecord{
// 		RequestedBy: c.GetString("uid"),
// 		RequestedAt: time.Now(),
// 	}
// 	err := service.UpdateRecord(UpdateRecordDto{Id: request.Id, Fullname: request.Fullname}, &audit)

// 	if err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	c.Status(http.StatusOK)
// }
