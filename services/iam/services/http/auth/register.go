package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	handlers "github.com/nezo32/sudoku/iam/services/handlers/auth"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func RegisterHandler(ctx echo.Context, entry *utils.EndpointEntry) error {
	first_name := ctx.FormValue("first_name")
	last_name := ctx.FormValue("last_name")
	username := ctx.FormValue("username")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	user := &model.Users{
		FirstName: &first_name,
		LastName:  &last_name,
		Username:  username,
		Email:     email,
		Password:  password,
	}

	res, err := handlers.RegisterUser(entry.ServiceContext, user)

	if err != nil {
		return err.ToHTTPError(ctx)
	}

	return ctx.JSON(http.StatusOK, errors.HTTPResponse{Data: res})
}
