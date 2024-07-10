package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"github.com/nezo32/sudoku/iam/services/http/utils"
)

func DefineUserEndpoints(echo *echo.Echo, database *pgxpool.Pool, ctx context.Context) {
	factory := utils.CreateEndpointFactory("/users", echo, database, ctx)

	factory.GET("/get", GetUserHanlder)
	factory.GET("/list", ListUserHanlder)
	factory.POST("/register", RegisterUserHandler)
}
