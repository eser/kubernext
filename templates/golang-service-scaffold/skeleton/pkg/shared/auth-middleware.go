package shared

import (
	"errors"
	"net/http"
	"strings"

	"github.com/eser/go-service/pkg/infra/httpserv"
)

type Token struct {
	UID   string `json:"uid"`
	Email string `json:"email"`
}

func VerifyIDToken[T interface{}](ctx *httpserv.Context, idToken string) (T, error) {
	// TODO(@eser): dummy implementation
	var token T

	// a sample failed verification
	// if len(token.UID) == 0 {
	//   err := errors.New("invalid user token")
	//   ctx.AbortWithError(http.StatusUnauthorized, err)

	//   return token, err
	// }

	return token, nil
}

func GetTokenFromHeader[T interface{}](ctx *httpserv.Context) (T, error) {
	var token T

	header := ctx.GetHeader("Authorization")

	idToken := strings.Split(header, " ")
	if len(idToken) < 2 || len(idToken[1]) == 0 {
		err := errors.New("invalid user token")
		ctx.AbortWithError(http.StatusUnauthorized, err)

		return token, err
	}

	verifiedToken, err := VerifyIDToken[T](ctx, idToken[1])
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return token, err
	}

	token = verifiedToken

	return token, nil
}

func AuthMiddleware() httpserv.HandlerFunc {
	return func(ctx *httpserv.Context) {
		// config, _ := config.GetConfig()

		verifiedToken, err := GetTokenFromHeader[Token](ctx)
		if err != nil {
			return
		}

		ctx.Set("uid", verifiedToken.UID)
		ctx.Set("email", verifiedToken.Email)
	}
}
