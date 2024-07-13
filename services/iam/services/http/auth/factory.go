package auth

import (
	"github.com/nezo32/sudoku/iam/services"
	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func DefineAuthEndpoints(ctx *services.ServiceContext) {
	group := ctx.Echo.Group("/auth")
	factory := utils.CreateEndpointFactory(ctx, group)

	factory.POST("/login", LoginHandler)
	factory.POST("/register", RegisterHandler)
	factory.POST("/refresh", RefreshHandler)
}
