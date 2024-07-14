package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/errors"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	"github.com/nezo32/sudoku/iam/services"
	handlers "github.com/nezo32/sudoku/iam/services/handlers/roles"
)

func RoleMiddleware(services_context *services.ServiceContext, role_titles ...model.RolesTitles) echo.MiddlewareFunc {
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

			neededRoleExists := checkRoles(role_titles, roles)

			if !neededRoleExists {
				ser_err := &errors.SerivceError{Code: http.StatusForbidden, Message: "doesn't have enough permissions"}
				return ser_err.ToHTTPError(ctx)
			}

			return next(ctx)
		}
	}
}

func checkRoles(role_constraints []model.RolesTitles, user_roles []model.Roles) bool {
	exists := false
	for _, role := range user_roles {
		for _, title := range role_constraints {
			switch title {
			case model.RolesTitles_IamAdmin:
				if role.Title == model.RolesTitles_IamAdmin {
					exists = true
				}
			case model.RolesTitles_IamEditor:
				if role.Title == model.RolesTitles_IamEditor || role.Title == model.RolesTitles_IamAdmin {
					exists = true
				}
			case model.RolesTitles_IamSelfEditor:
				if role.Title == model.RolesTitles_IamEditor || role.Title == model.RolesTitles_IamAdmin || role.Title == model.RolesTitles_IamSelfEditor {
					exists = true
				}
			case model.RolesTitles_IamViewer:
				if role.Title == model.RolesTitles_IamEditor || role.Title == model.RolesTitles_IamAdmin || role.Title == model.RolesTitles_IamViewer {
					exists = true
				}
			}
		}
		if exists {
			break
		}
	}
	return exists
}
