package roles

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/services/handlers/roles"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func GetUserRolesByIDHandler(ctx echo.Context, entry *utils.EndpointEntry) error {
	input := ctx.FormValue("id")

	id, err := uuid.Parse(input)
	if err != nil {
		ser_err := &errors.SerivceError{Code: http.StatusBadRequest, Error: err, Message: "provided id is not valid uuid"}
		return ser_err.ToHTTPError(ctx)
	}

	roles, ser_err := roles.GetUserRolesByID(entry.ServiceContext, &id)
	if ser_err != nil {
		return ser_err.ToHTTPError(ctx)
	}

	return ctx.JSON(http.StatusOK, errors.HTTPResponse{Data: roles})
}
