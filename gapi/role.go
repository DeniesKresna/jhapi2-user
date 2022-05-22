package gapi

import (
	"context"

	pbRole "github.com/DeniesKresna/jhapi2-user/pb/role"
	"github.com/DeniesKresna/jhapi2-user/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateRole(ctx context.Context, req *pbRole.CreateRoleRequest) (*pbRole.CreateRoleResponse, error) {
	createRoleRequest := types.CreateRoleRequest{}
	server.utils.InjectStructValue(req, &createRoleRequest)

	role, err := server.roleHandler.RoleCreate(ctx, createRoleRequest)
	if err != nil {
		return nil, status.Errorf(codes.Canceled, err.Error())
	}

	userResponse := &pbRole.CreateRoleResponse{
		Role: types.ConvertRole(role),
	}
	return userResponse, nil
}
