package roles

import (
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/middlewares"
	"github.com/nezo32/sudoku/iam/services"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func DefineRolesEndpoints(ctx *services.ServiceContext) {
	group := ctx.Echo.Group("/roles")
	factory := utils.CreateEndpointFactory(ctx, group)

	factory.GET("/get_by_id", GetUserRolesByIDHandler, middlewares.RoleMiddleware(ctx, model.RolesTitles_IamViewer))
}
