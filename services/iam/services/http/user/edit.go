package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/services"
	"github.com/nezo32/sudoku/iam/services/handlers/user"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func EditUserHandler(ctx echo.Context, entry *utils.EndpointEntry) error {
	input := &services.UserEditInput{}
	err := ctx.Bind(input)
	if err != nil {
		ser_err := &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
		return ser_err.ToHTTPError(ctx)
	}

	user, ser_err := user.EditUserHandler(entry.ServiceContext, input)
	if ser_err != nil {
		return ser_err.ToHTTPError(ctx)
	}

	return ctx.JSON(http.StatusOK, errors.HTTPResponse{Data: user})
}
