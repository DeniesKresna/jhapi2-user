package handlers

import (
	"context"

	"github.com/DeniesKresna/jhapi2-user/repos"
	"github.com/DeniesKresna/jhapi2-user/types"
)

type UserHandler struct {
	userRepo repos.IUserRepository
}

func ProvideUserHandler(userRepo repos.IUserRepository) IUserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (h UserHandler) UserCreate(ctx context.Context, req types.CreateUserRequest) (u types.User, err error) {
	if u, err = h.userRepo.UserCreate(ctx, req); err != nil {
		return u, err
	}
	return u, nil
}
