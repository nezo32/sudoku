package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/services/handlers/user"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func GetUserHanlder(ctx echo.Context, entry *utils.EndpointEntry) error {
	input := ctx.FormValue("id")

	user, err := user.GetUserByID(entry.ServiceContext, input)
	if err != nil {
		return err.ToHTTPError(ctx)
	}

	return ctx.JSON(http.StatusOK, errors.HTTPResponse{Data: user})
}
