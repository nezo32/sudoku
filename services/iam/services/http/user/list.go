package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/services/handlers/user"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func ListUserHanlder(ctx echo.Context, entry *utils.EndpointEntry) error {
	users, err := user.ListUsers(entry.ServiceContext)

	if err != nil {
		return err.ToHTTPError(ctx)
	}

	return ctx.JSON(http.StatusOK, errors.HTTPResponse{Data: users})
}
