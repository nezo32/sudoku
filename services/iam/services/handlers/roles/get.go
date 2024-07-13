package roles

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/gommon/log"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/services"
)

func GetUserRolesByID(ctx *services.ServiceContext, input *uuid.UUID) ([]model.Roles, *errors.SerivceError) {
	rows, err := ctx.Database.Query(ctx.Context, `select roles.id, roles.title from users, user_roles, roles
	 where users.id = $1 and users.id = user_roles.user_id and roles.id = user_roles.role_id`, input.String())
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Roles])
	if err != nil {
		log.Error(err)
		if pgx.ErrNoRows == err {
			return res, nil
		}
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}

	return res, nil
}
