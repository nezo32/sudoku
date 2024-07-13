package user

import (
	"github.com/nezo32/sudoku/iam/services"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func DefineUserEndpoints(ctx *services.ServiceContext) {
	factory := utils.CreateEndpointFactory("/users", ctx)

	factory.GET("/get", GetUserHanlder)
	factory.GET("/list", ListUserHanlder)
}
