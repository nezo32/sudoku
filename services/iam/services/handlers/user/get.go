package user

import (
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/services"
)

func GetUserByID(ctx services.ServiceContext, input string) (*model.Users, error) {
	id, err := uuid.Parse(input)
	if err != nil {
		return nil, err
	}

	user := model.Users{}
	rows, err := ctx.Database.Query(ctx.Context, `select first_name, last_name, email, created_at, updated_at from users where id=$1`, id)
	if err != nil {
		return nil, err
	}
	err = pgxscan.ScanOne(&user, rows)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
