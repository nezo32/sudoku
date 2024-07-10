package permissions_service

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	permpb "github.com/nezo32/sudoku/iam/generated/protos/permissions"
	"github.com/nezo32/sudoku/iam/services"
	"github.com/nezo32/sudoku/iam/services/http/permissions"
)

type PermissionsServiceServer struct {
	services.ServiceContext
	permpb.UnimplementedPermissionsServiceServer
}

func CreateUserSerivceServer(e *echo.Echo, db *pg.DB) *PermissionsServiceServer {
	permissions.DefinePermissonsEndpoints(e, db)
	return &PermissionsServiceServer{ServiceContext: services.ServiceContext{Echo: e, Database: db}}
}

func (server *PermissionsServiceServer) Get(ctx context.Context, req *permpb.PermissionsRequest) (*permpb.PermissionsResponse, error) {
	return nil, errors.New("unimplemented")
}