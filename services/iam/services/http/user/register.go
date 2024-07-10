package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/services"
	handlers "github.com/nezo32/sudoku/iam/services/handlers/user"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func RegisterUserHandler(ctx echo.Context, entry *utils.EndpointEntry) error {
	first_name := ctx.FormValue("first_name")
	last_name := ctx.FormValue("last_name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	user := model.Users{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Password:  password,
	}

	err := handlers.RegisterUser(services.ServiceContext{Database: entry.Database}, user)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}
