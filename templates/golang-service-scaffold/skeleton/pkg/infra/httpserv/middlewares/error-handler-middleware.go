package middlewares

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
	if fe.Tag() == "required" {
		return fmt.Sprintf("Field %s is required", fe.Field())
	}

	return "Unknown error"
}

func ErrorHandlerMiddleware() httpserv.HandlerFunc {
	return func(ctx *httpserv.Context) {
		ctx.Next()

		var (
			out                    []ErrorMsg
			hasNonValidationErrors bool = false
		)

		for _, err := range ctx.Errors {
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

			hasNonValidationErrors = true

			out = append(out, ErrorMsg{err.Error()})
		}

		if len(out) > 0 {
			json := httpserv.H{"errors": out}

			if hasNonValidationErrors {
				// status -1 doesn't overwrite existing status code
				// ctx.JSON(-1, json)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, json)
			} else {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, json)
			}
		}
	}
}
