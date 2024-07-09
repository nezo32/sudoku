package user

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"

	"github.com/nezo32/sudoku/iam/endpoints/utils"
)

func DefineUserEndpoints(echo *echo.Echo, database *pg.DB) {
	factory := utils.CreateEndpointFactory("/users", echo, database)

	factory.GET("/get", GetUserHanlder)
	factory.GET("/list", ListUserHanlder)
}
