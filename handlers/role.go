package handlers

import (
	"context"

	"github.com/DeniesKresna/jhapi2-user/repos"
	"github.com/DeniesKresna/jhapi2-user/types"
)

type RoleHandler struct {
	roleRepo repos.IRoleRepository
}

func ProvideRoleHandler(roleRepo repos.IRoleRepository) IRoleHandler {
	return &RoleHandler{
		roleRepo: roleRepo,
	}
}

func (h RoleHandler) RoleCreate(ctx context.Context, req types.CreateRoleRequest) (role types.Role, err error) {
	if role, err = h.roleRepo.RoleCreate(ctx, req); err != nil {
		return role, err
	}
	return role, nil
}
