package auth

import (
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/services"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func RegisterUser(ctx *services.ServiceContext, input *model.Users) (*string, *errors.SerivceError) {

	if input.Username == "" {
		return nil, &errors.SerivceError{Code: http.StatusBadRequest, Message: "username can't be empty"}
	}

	if input.Email == "" {
		return nil, &errors.SerivceError{Code: http.StatusBadRequest, Message: "email can't be empty"}
	}

	if !utils.IsEmailValid(input.Email) {
		return nil, &errors.SerivceError{Code: http.StatusBadRequest, Message: "invalid email address"}
	}

	if input.Password == "" {
		return nil, &errors.SerivceError{Code: http.StatusBadRequest, Message: "password can't be empty"}
	}

	if len(input.Password) < 8 {
		return nil, &errors.SerivceError{Code: http.StatusBadRequest, Message: "password is too short"}
	}

	if len(input.Username) < 4 {
		return nil, &errors.SerivceError{Code: http.StatusBadRequest, Message: "username is too short"}
	}

	password, err := ctx.PasswordEncode(input.Password)

	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Message: "can't generate password hash", Error: err}
	}

	tx, err := ctx.Database.Begin(ctx.Context)
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}

	var id string
	err = tx.QueryRow(ctx.Context, `insert into users (username, first_name, last_name, email, password) values ($1, $2, $3, $4, $5) returning id`,
		input.Username, input.FirstName, input.LastName, input.Email, password).Scan(&id)
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}
	_, err = tx.Exec(ctx.Context, `insert into user_roles (user_id, role_id) values ($1, (select id from roles where title='iam.viewer'))`, id)
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}
	err = tx.Commit(ctx.Context)
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}

	return &id, nil
}
