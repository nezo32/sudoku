package user

import (
	"github.com/google/uuid"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/services"
)

func GetUserByID(ctx services.ServiceContext, input string) (*model.Users, error) {
	id, err := uuid.Parse(input)
	if err != nil {
		return nil, err
	}

	user := &model.Users{ID: id}
	err = ctx.Database.Model(user).WherePK().Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}
