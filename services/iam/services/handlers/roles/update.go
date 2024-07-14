package roles

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/gommon/log"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/services"
)

func UpdateUserRoles(ctx *services.ServiceContext, id *uuid.UUID, roles ...model.RolesTitles) ([]model.Roles, *errors.SerivceError) {
	if len(roles) == 0 {
		return nil, &errors.SerivceError{Code: http.StatusBadRequest, Message: "nothing to update"}
	}

	tx, err := ctx.Database.Begin(ctx.Context)
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}

	insert_args := []any{id.String()}
	var sqlSelects []string

	for i, role := range roles {
		insert_args = append(insert_args, role.String())
		sqlSelects = append(sqlSelects, fmt.Sprintf("($1, (select id from roles where title=$%d))", i+2))
	}

	_, err = tx.Exec(ctx.Context, `insert into user_roles (user_id, role_id) values`+strings.Join(sqlSelects, ","), insert_args...)
	if err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusBadRequest, Error: err}
	}

	rows, err := tx.Query(ctx.Context, `select roles.id, roles.title from users, user_roles, roles
	 where users.id = $1 and users.id = user_roles.user_id and roles.id = user_roles.role_id`, id.String())
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

	if err = tx.Commit(ctx.Context); err != nil {
		log.Error(err)
		return nil, &errors.SerivceError{Code: http.StatusInternalServerError, Error: err}
	}

	return res, nil
}
