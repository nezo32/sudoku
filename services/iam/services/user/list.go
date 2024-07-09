package user

import (
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/services"
)

func ListUsers(ctx services.ServiceContext) ([]model.Users, error) {
	var users []model.Users
	err := ctx.Database.Model(&users).Column("id", "first_name", "last_name", "email", "created_at", "updated_at").Select()
	if err != nil {
		return nil, err
	}

	return users, nil
}
