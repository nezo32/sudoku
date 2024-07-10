package user

import (
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/services"
)

func ListUsers(ctx services.ServiceContext) ([]*model.Users, error) {
	var users []*model.Users
	err := pgxscan.Select(ctx.Context, ctx.Database, &users, `select first_name, last_name, email, created_at, updated_at from users`)
	if err != nil {
		return nil, err
	}

	return users, nil
}
