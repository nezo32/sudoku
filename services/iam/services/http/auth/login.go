package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/services/handlers/auth"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func LoginHandler(ctx echo.Context, entry *utils.EndpointEntry) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	token, err := auth.Login(entry.ServiceContext, &auth.LoginInput{
		Username: username,
		Password: password,
	})

	if err != nil {
		return err.ToHTTPError(ctx)
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
