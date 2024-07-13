package user_service

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/nezo32/sudoku/iam/generated/postgres/IAM/public/model"
	userpb "github.com/nezo32/sudoku/iam/generated/protos/user"
	"github.com/nezo32/sudoku/iam/services"
	handlers "github.com/nezo32/sudoku/iam/services/handlers/user"
	"github.com/nezo32/sudoku/iam/services/http/user"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserSerivceServer struct {
	*services.ServiceContext
	userpb.UnimplementedUserServiceServer
}

func CreateUserSerivceServer(ctx *services.ServiceContext) *UserSerivceServer {
	user.DefineUserEndpoints(ctx)
	return &UserSerivceServer{ServiceContext: ctx}
}

func (server *UserSerivceServer) Get(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	res, err := handlers.GetUserByID(server.ServiceContext, req.Id)

	if err != nil {
		log.Error(err)
		log.Error(err.Error)
		return nil, err.ToGRPCError()
	}

	return dbToProtoUser(&model.Users{
		ID:        res.ID,
		Username:  res.Username,
		FirstName: &res.FirstName,
		LastName:  &res.LastName,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}), nil
}

func (server *UserSerivceServer) List(ctx context.Context, _ *emptypb.Empty) (*userpb.UserListResponse, error) {
	res, err := handlers.ListUsers(server.ServiceContext)

	if err != nil {
		log.Error(err)
		log.Error(err.Error)
		return nil, err.ToGRPCError()
	}

	var users []*userpb.UserResponse

	for _, user := range res {

		var UpdatedAt *timestamppb.Timestamp = nil
		if user.UpdatedAt != nil {
			UpdatedAt = timestamppb.New(*user.UpdatedAt)
		}
		users = append(users, &userpb.UserResponse{
			Id:        user.ID.String(),
			Username:  user.Username,
			FirstName: &user.FirstName,
			LastName:  &user.LastName,
			Email:     user.Email,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: UpdatedAt,
		})
	}

	return &userpb.UserListResponse{
		UserList: users,
	}, nil
}

func dbToProtoUser(input *model.Users) *userpb.UserResponse {
	var UpdatedAt *timestamppb.Timestamp = nil
	if input.UpdatedAt != nil {
		UpdatedAt = timestamppb.New(*input.UpdatedAt)
	}

	return &userpb.UserResponse{
		Id:        input.ID.String(),
		Username:  input.Username,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		CreatedAt: timestamppb.New(input.CreatedAt),
		UpdatedAt: UpdatedAt,
	}
}
