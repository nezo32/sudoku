package middlewares

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/security"
)

func jwtCheckMiddleware(ctx echo.Context) (*security.TokenSub, *errors.SerivceError) {
	token := ""
	access_token_cookie, _ := ctx.Cookie("access_token")

	if access_token_cookie != nil {
		token = access_token_cookie.Value
	} else {
		token = ctx.Request().Header.Get("Authorization")
		if token == "" {
			return nil, &errors.SerivceError{Code: http.StatusUnauthorized, Message: "token is not provided"}
		}
		_, err := fmt.Sscanf(token, "Bearer %s", &token)
		if err != nil {
			return nil, &errors.SerivceError{Code: http.StatusUnauthorized,
				Error: err, Message: "token doesn't satisfy format constraint"}
		}
	}

	if token == "" {
		return nil, &errors.SerivceError{Code: http.StatusUnauthorized, Message: "token is not provided"}
	}

	user, err := security.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			_, err := jwtCheckMiddleware(ctx)

			if err != nil {
				return err.ToHTTPError(ctx)
			}

			return next(ctx)
		}
	}
}
