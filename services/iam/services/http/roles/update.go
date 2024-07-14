package roles

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/services/handlers/roles"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

type UpdateUserRolesInput struct {
	ID    string   `json:"id,omitempty" form:"id"`
	Roles []string `json:"roles,omitempty" form:"roles"`
}

func UpdateUserRolesHandler(ctx echo.Context, entry *utils.EndpointEntry) error {
	input := &UpdateUserRolesInput{}
	err := ctx.Bind(input)
	if err != nil {
		ser_err := &errors.SerivceError{Code: http.StatusBadRequest, Error: err}
		return ser_err.ToHTTPError(ctx)
	}

	if input.ID == "" || len(input.Roles) == 0 {
		ser_err := &errors.SerivceError{Code: http.StatusBadRequest, Message: "id and roles can't be empty"}
		return ser_err.ToHTTPError(ctx)
	}

	id, err := uuid.Parse(input.ID)
	if err != nil {
		ser_err := &errors.SerivceError{Code: http.StatusBadRequest, Error: err, Message: "provided id is not valid uuid"}
		return ser_err.ToHTTPError(ctx)
	}

	input_roles := []model.RolesTitles{}

	for _, role := range input.Roles {
		tmp := model.RolesTitles_IamViewer
		err = (&tmp).Scan(role)
		if err != nil {
			ser_err := &errors.SerivceError{Code: http.StatusBadRequest, Error: err, Message: fmt.Sprintf("role - %s doesn't exists", role)}
			return ser_err.ToHTTPError(ctx)
		}
		input_roles = append(input_roles, tmp)
	}

	roles, ser_err := roles.UpdateUserRoles(entry.ServiceContext, &id, input_roles...)
	if ser_err != nil {
		return ser_err.ToHTTPError(ctx)
	}

	return ctx.JSON(http.StatusOK, errors.HTTPResponse{Data: roles})
}
