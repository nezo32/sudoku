package middlewares

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/security"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token := ""
			access_token_cookie, _ := ctx.Cookie("access_token")

			if access_token_cookie != nil {
				token = access_token_cookie.Value
			} else {
				token = ctx.Request().Header.Get("Authorization")
				_, err := fmt.Sscanf(token, "Bearer %s", &token)
				if err != nil {
					return ctx.JSON(http.StatusUnauthorized, errors.HTTPResponse{Error: &errors.SerivceError{Code: http.StatusUnauthorized,
						Error: err, Message: "token doesn't satisfy format constraint"}})
				}
			}

			if token == "" {
				return ctx.JSON(http.StatusUnauthorized, errors.HTTPResponse{Error: &errors.SerivceError{Code: http.StatusUnauthorized, Message: "token is not provided"}})
			}

			_, err := security.ValidateToken(token)
			if err != nil {
				return ctx.JSON(http.StatusUnauthorized, errors.HTTPResponse{Error: err})
			}

			return next(ctx)
		}
	}
}
