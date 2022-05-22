package handlers

import (
	"context"

	"github.com/DeniesKresna/jhapi2-user/types"
)

type IUserHandler interface {
	UserCreate(ctx context.Context, req types.CreateUserRequest) (u types.User, err error)
}
