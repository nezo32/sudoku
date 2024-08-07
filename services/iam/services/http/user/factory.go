package user

import (
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/middlewares"
	"github.com/nezo32/sudoku/iam/services"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func DefineUserEndpoints(ctx *services.ServiceContext) {
	group := ctx.Echo.Group("/users")
	factory := utils.CreateEndpointFactory(ctx, group)

	factory.GET("/get", GetUserHanlder)
	factory.GET("/list", ListUserHanlder)
	factory.PUT("/edit", EditUserHandler, middlewares.RoleMiddleware(ctx, model.RolesTitles_IamEditor))
}
