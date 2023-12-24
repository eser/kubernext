package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/eser/go-service/pkg/infra/httpserv"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrMissingUserToken = errors.New("missing user token")
	ErrInvalidUserToken = errors.New("invalid user token")
)

type AuthInformation struct {
	jwt.RegisteredClaims
	UID   string `json:"uid"`
	Email string `json:"email"`
}

func VerifyIdToken(jwtSignature string, idToken string) (*AuthInformation, error) {
	if len(idToken) == 0 {
		return nil, ErrMissingUserToken
	}

	token, err := jwt.ParseWithClaims(idToken, &AuthInformation{}, func(tokenPreview *jwt.Token) (any, error) {
		// validate the alg is what you expect:
		if _, ok := tokenPreview.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v: %w", tokenPreview.Header["alg"], ErrInvalidUserToken)
		}

		secret := []byte(jwtSignature)

		return secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error parsing token: %w", err)
	}

	if token.Valid {
		if claims, ok := token.Claims.(*AuthInformation); ok {
			return claims, nil
		}
	}

	return nil, fmt.Errorf("Token is invalid: %w", ErrInvalidUserToken)
}

func ReadTokenFromString(jwtSignature, header string) (*AuthInformation, error) {
	idToken := strings.SplitN(header, " ", 2) //nolint:gomnd

	if len(idToken) < 2 || len(idToken[1]) == 0 {
		return nil, ErrInvalidUserToken
	}

	verifiedToken, err := VerifyIdToken(jwtSignature, idToken[1])
	if err != nil {
		return nil, err
	}

	return verifiedToken, nil
}

func AuthMiddleware(jwtSignature string, isRequired bool, isDevelopmentEnv bool) httpserv.HandlerFunc {
	return func(ctx *httpserv.Context) {
		if isDevelopmentEnv {
			ctx.Set("uid", "01HFVDN8AY8NJYAHJ3MD1MQJPZ") // Eser's Static Person ID
			ctx.Set("email", "eser@acikyazilim.com")     // Eser's Static Person Email

			ctx.Next()

			return
		}

		// config, _ := config.GetConfig()
		header := ctx.GetHeader("Authorization")

		if len(header) == 0 {
			if isRequired {
				ctx.AbortWithError(http.StatusUnauthorized, ErrMissingUserToken) //nolint:errcheck

				return
			}

			ctx.Next()

			return
		}

		verifiedToken, err := ReadTokenFromString(jwtSignature, header)
		if err != nil {
			return
		}

		ctx.Set("uid", verifiedToken.UID)
		ctx.Set("email", verifiedToken.Email)

		ctx.Next()
	}
}
