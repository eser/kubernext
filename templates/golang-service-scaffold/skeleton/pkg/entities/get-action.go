package entities

// import (
// 	"net/http"
// 	"time"

// 	"github.com/eser/go-service/pkg/infra/httpserv"
// 	"github.com/eser/go-service/pkg/shared"
// )

// type GetActionRequest struct {
// 	Id string `json:"id"`
// }

// type GetActionResponse struct {
// 	Record *Model `json:"record"`
// }

// func GetAction(c *httpserv.Context) {
// 	request := new(GetActionRequest)

// 	if err := c.ShouldBind(&request); err != nil {
// 		c.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	service := NewService()
// 	audit := shared.AuditRecord{
// 		RequestedBy: c.GetString("uid"),
// 		RequestedAt: time.Now(),
// 	}
// 	record, err := service.GetRecord(request.Id, &audit)

// 	if err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, GetActionResponse{record})
// }
