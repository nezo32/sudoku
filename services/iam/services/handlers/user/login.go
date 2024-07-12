package user

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/gommon/log"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/security"
	"github.com/nezo32/sudoku/iam/services"
)

type LoginInput struct {
	Username string
	Password string
}

func Login(ctx *services.ServiceContext, input *LoginInput) (*security.TokenPair, *errors.SerivceError) {
	rows, err := ctx.Database.Query(ctx.Context, `select * from users where username=$1`, input.Username)
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}
	res, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Users])
	if err != nil {
		log.Error(err)
		if pgx.ErrNoRows == err {
			return nil, &errors.SerivceError{Code: http.StatusBadRequest, Message: "user/password is incorrect"}
		}
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}

	passwordIsCorrect, err := security.CheckPasswordHash(input.Password, res.Password)
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}
	if !passwordIsCorrect {
		return nil, &errors.SerivceError{Code: http.StatusBadRequest, Message: "user/password is incorrect"}
	}

	token, sErr := ctx.JWTGenerator(security.TokenSub{ID: res.ID, Username: res.Username, Email: res.Email})

	if sErr != nil {
		return nil, sErr
	}

	return token, nil
}
