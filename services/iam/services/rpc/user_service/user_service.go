package user_service

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	userpb "github.com/nezo32/sudoku/iam/generated/protos/user"
	"github.com/nezo32/sudoku/iam/services"
	handlers "github.com/nezo32/sudoku/iam/services/handlers/user"
	"github.com/nezo32/sudoku/iam/services/http/user"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserSerivceServer struct {
	services.ServiceContext
	userpb.UnimplementedUserServiceServer
}

func CreateUserSerivceServer(e *echo.Echo, db *pg.DB) *UserSerivceServer {
	user.DefineUserEndpoints(e, db)
	return &UserSerivceServer{ServiceContext: services.ServiceContext{Echo: e, Database: db}}
}

func (server *UserSerivceServer) Get(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	res, err := handlers.GetUserByID(server.ServiceContext, req.Id)

	if err != nil {
		return nil, err
	}

	return dbToProtoUser(res), nil
}

func (server *UserSerivceServer) List(ctx context.Context, _ *emptypb.Empty) (*userpb.UserListResponse, error) {
	res, err := handlers.ListUsers(server.ServiceContext)

	if err != nil {
		return nil, err
	}

	var users []*userpb.UserResponse

	for _, user := range res {
		users = append(users, dbToProtoUser(&user))
	}

	return &userpb.UserListResponse{
		UserList: users,
	}, nil
}

func dbToProtoUser(input *model.Users) *userpb.UserResponse {
	var UpdatedAt *timestamppb.Timestamp = nil
	if input.UpdatedAt != nil {
		UpdatedAt = timestamppb.New(input.CreatedAt)
	}

	return &userpb.UserResponse{
		Id:        input.ID.String(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		CreatedAt: timestamppb.New(input.CreatedAt),
		UpdatedAt: UpdatedAt,
	}
}
