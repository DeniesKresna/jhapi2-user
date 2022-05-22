package repos

import (
	"context"

	"github.com/DeniesKresna/jhapi2-user/types"
)

type IRoleRepository interface {
	RoleCreate(ctx context.Context, req types.CreateRoleRequest) (u types.Role, err error)
	RoleGetByID(ctx context.Context, id int64) (u types.Role, err error)
}
