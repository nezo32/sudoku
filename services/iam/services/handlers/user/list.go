package user

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/gommon/log"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/services"
)

func ListUsers(ctx *services.ServiceContext) ([]services.UsersOutput, *errors.SerivceError) {
	rows, err := ctx.Database.Query(ctx.Context, `select id, username, first_name, last_name, email, created_at, updated_at from users`)
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[services.UsersOutput])
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Message: "can't list users",
			Error: err}
	}
	return users, nil
}
