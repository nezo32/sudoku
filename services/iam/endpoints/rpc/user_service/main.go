package user_service

import (
	"context"
	"errors"

	userpb "github.com/nezo32/sudoku/iam/generated/protos/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserSerivceServer struct {
	userpb.UnimplementedUserServiceServer
}

func CreateUserSerivceServer() *UserSerivceServer {
	return &UserSerivceServer{}
}

func (server *UserSerivceServer) Get(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	return nil, errors.New("unimplemented")
}

func (server *UserSerivceServer) List(ctx context.Context, _ *emptypb.Empty) (*userpb.UserListResponse, error) {
	return nil, errors.New("unimplemented")
}
