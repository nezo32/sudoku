package user

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/gommon/log"

	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/services"
)

func GetUserByID(ctx *services.ServiceContext, input string) (*services.UsersOutput, *errors.SerivceError) {
	id, err := uuid.Parse(input)
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusBadRequest, Message: "invalid uuid", Error: err}
	}

	rows, err := ctx.Database.Query(ctx.Context, `select id, username, first_name, last_name, email, created_at, updated_at from users where id=$1`, id)
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}
	res, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[services.UsersOutput])
	if err != nil {
		log.Error(err)
		if pgx.ErrNoRows == err {
			return nil, &errors.SerivceError{Code: http.StatusNotFound, Message: "user not found", Error: err}
		}
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}

	return res, nil
}
