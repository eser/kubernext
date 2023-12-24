package shared

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/eser/go-service/pkg/infra/httpserv"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Message string `json:"message"`
}

func getErrorMessageFromFieldError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("Field %s is required", fe.Field())
	}

	return "Unknown error"
}

func ErrorHandlerMiddleware() httpserv.HandlerFunc {
	return func(c *httpserv.Context) {
		c.Next()

		var out []ErrorMsg

		for _, err := range c.Errors {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				for _, fe := range ve {
					out = append(
						out,
						ErrorMsg{getErrorMessageFromFieldError(fe)},
					)
				}

				continue
			}

			out = append(
				out,
				ErrorMsg{err.Error()},
			)
		}

		if len(out) > 0 {
			json := httpserv.H{"errors": out}

			c.AbortWithStatusJSON(http.StatusInternalServerError, json)
			// status -1 doesn't overwrite existing status code
			// c.JSON(-1, json)
		}
	}
}
