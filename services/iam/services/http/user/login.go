package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/services/handlers/user"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func LoginHandler(ctx echo.Context, entry *utils.EndpointEntry) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	token, err := user.Login(entry.ServiceContext, &user.LoginInput{
		Username: username,
		Password: password,
	})

	if err != nil {
		return err.ToHTTPError(ctx)
	}

	ctx.SetCookie(&http.Cookie{
		Name:  "token",
		Value: token.AccessToken,
	})

	return ctx.JSON(http.StatusOK, errors.HTTPResponse{Data: token})
}
