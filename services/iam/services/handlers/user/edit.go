package user

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/gommon/log"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/services"
)

func EditUserHandler(ctx *services.ServiceContext, input *services.UserEditInput) (*services.UsersOutput, *errors.SerivceError) {
	if input.ID == "" {
		return nil, &errors.SerivceError{Code: http.StatusBadRequest, Message: "user id is not provided"}
	}

	tx, err := ctx.Database.Begin(ctx.Context)
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Message: "can't start transaction", Error: err}
	}

	if input.FirstName != "" {
		_, err = tx.Exec(ctx.Context, `update users set first_name = $2 where id = $1`, input.ID, input.FirstName)
		if err != nil {
			log.Error(err)
			return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
		}
	}
	if input.LastName != "" {
		_, err = tx.Exec(ctx.Context, `update users set last_name = $2 where id = $1`, input.ID, input.LastName)
		if err != nil {
			log.Error(err)
			return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
		}
	}
	if input.Password != "" {
		_, err = tx.Exec(ctx.Context, `update users set password = $2 where id = $1`, input.ID, input.Password)
		if err != nil {
			log.Error(err)
			return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
		}
	}
	if input.Username != "" {
		_, err = tx.Exec(ctx.Context, `update users set username = $2 where id = $1`, input.ID, input.Username)
		if err != nil {
			log.Error(err)
			return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
		}
	}

	rows, err := tx.Query(ctx.Context, `select id, username, first_name, last_name, email, created_at, updated_at from users where id=$1`, input.ID)
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

	if err = tx.Commit(ctx.Context); err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Message: "can't commit transaction", Error: err}
	}

	return res, nil
}
