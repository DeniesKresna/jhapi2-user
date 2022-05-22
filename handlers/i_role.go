package handlers

import (
	"context"

	"github.com/DeniesKresna/jhapi2-user/types"
)

type IRoleHandler interface {
	RoleCreate(ctx context.Context, req types.CreateRoleRequest) (role types.Role, err error)
}
