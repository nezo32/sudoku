package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/endpoints/utils"
	"github.com/nezo32/sudoku/iam/services"
	"github.com/nezo32/sudoku/iam/services/user"
)

func ListUserHanlder(ctx echo.Context, entry *utils.EndpointEntry) error {
	users, err := user.ListUsers(services.ServiceContext{Database: entry.Database})

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, users)
}
