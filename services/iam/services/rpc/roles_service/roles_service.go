package roles_service

import (
	"context"
	"errors"

	rolespb "github.com/nezo32/sudoku/iam/generated/protos/roles"
	"github.com/nezo32/sudoku/iam/services"
	"github.com/nezo32/sudoku/iam/services/http/roles"
)

type RolesServiceServer struct {
	*services.ServiceContext
	rolespb.UnimplementedRolesServiceServer
}

func CreateRolesServiceServer(ctx *services.ServiceContext) *RolesServiceServer {
	roles.DefineRolesEndpoints(ctx)
	return &RolesServiceServer{ServiceContext: ctx}
}

func (server *RolesServiceServer) GetByUserID(ctx context.Context, req *rolespb.RolesRequest) (*rolespb.RolesResponse, error) {
	return nil, errors.New("unimplemented")
}
