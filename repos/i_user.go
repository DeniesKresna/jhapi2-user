package repos

import (
	"context"

	"github.com/DeniesKresna/jhapi2-user/types"
)

type IUserRepository interface {
	UserCreate(ctx context.Context, req types.CreateUserRequest) (u types.User, err error)
	UserGetByID(ctx context.Context, id int64) (u types.User, err error)
}
