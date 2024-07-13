package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/security"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func RefreshHandler(ctx echo.Context, entry *utils.EndpointEntry) error {
	token := &security.TokenPair{}

	access_token_cookie, _ := ctx.Cookie("access_token")
	refresh_token_cookie, _ := ctx.Cookie("refresh_token")

	if access_token_cookie != nil && refresh_token_cookie != nil {
		token.AccessToken = access_token_cookie.Value
		token.RefreshToken = refresh_token_cookie.Value
	} else {
		err := ctx.Bind(token)
		if err != nil {
			ser_err := &errors.SerivceError{Code: http.StatusBadRequest, Error: err}
			return ser_err.ToHTTPError(ctx)
		}
	}

	if token.AccessToken == "" || token.RefreshToken == "" {
		return ctx.JSON(http.StatusBadRequest, errors.HTTPResponse{Error: &errors.SerivceError{Code: http.StatusBadRequest, Message: "tokens are not provided"}})
	}

	user, ser_err := security.ValidateRefreshToken(token)
	if ser_err != nil {
		return ser_err.ToHTTPError(ctx)
	}

	token, ser_err = entry.JWTGenerator(user)
	if ser_err != nil {
		return ser_err.ToHTTPError(ctx)
	}

	ctx.SetCookie(&http.Cookie{
		Name:  "access_token",
		Value: token.AccessToken,
		Path:  "/",
	})
	ctx.SetCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: token.RefreshToken,
		Path:  "/",
	})

	return ctx.JSON(http.StatusOK, errors.HTTPResponse{Data: token})
}
