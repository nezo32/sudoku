package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/services"
	handlers "github.com/nezo32/sudoku/iam/services/handlers/roles"
)

func RoleMiddleware(services_context *services.ServiceContext, roleTitles ...model.RolesTitles) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			user, err := jwtCheckMiddleware(ctx)

			if err != nil {
				return err.ToHTTPError(ctx)
			}

			if user == nil {
				ser_err := &errors.SerivceError{Code: http.StatusUnauthorized, Message: "user info not found"}
				return ser_err.ToHTTPError(ctx)
			}

			roles, err := handlers.GetUserRolesByID(services_context, &user.ID)
			if err != nil {
				return err.ToHTTPError(ctx)
			}

			var neededRoleExists bool
			for _, role := range roles {
				for _, title := range roleTitles {
					if role.Title == title {
						neededRoleExists = true
						break
					}
				}
			}

			if !neededRoleExists {
				ser_err := &errors.SerivceError{Code: http.StatusForbidden, Message: "doesn't have enough permissions"}
				return ser_err.ToHTTPError(ctx)
			}

			return next(ctx)
		}
	}
}
