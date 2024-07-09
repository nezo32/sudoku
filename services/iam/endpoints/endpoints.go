package endpoints

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"

	"github.com/nezo32/sudoku/iam/endpoints/permissions"
	"github.com/nezo32/sudoku/iam/endpoints/user"
)

func DefineEndpoints(e *echo.Echo, db *pg.DB) {
	user.DefineUserEndpoints(e, db)
	permissions.DefinePermissonsEndpoints(e)
}
