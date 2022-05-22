package gapi

import (
	"context"

	pbUser "github.com/DeniesKresna/jhapi2-user/pb/user"
	"github.com/DeniesKresna/jhapi2-user/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pbUser.CreateUserRequest) (*pbUser.CreateUserResponse, error) {
	createUserRequest := types.CreateUserRequest{}
	server.utils.InjectStructValue(req, &createUserRequest)

	user, err := server.userHandler.UserCreate(ctx, createUserRequest)
	if err != nil {
		return nil, status.Errorf(codes.Canceled, err.Error())
	}

	userResponse := &pbUser.CreateUserResponse{
		User: types.ConvertUser(user),
	}
	return userResponse, nil
}
