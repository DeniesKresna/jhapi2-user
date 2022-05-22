package repos

import (
	"context"
	"fmt"

	"github.com/DeniesKresna/jhapi2-user/types"
	"github.com/DeniesKresna/jhapi2/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	db    *gorm.DB
	utils utils.IUtils
}

func ProvideUserRepository(db *gorm.DB, utils utils.IUtils) IUserRepository {
	return &UserRepository{
		db:    db,
		utils: utils,
	}
}

func (r UserRepository) UserGetByID(ctx context.Context, id int64) (u types.User, err error) {
	err = r.db.First(&u, id).Error
	return u, err
}

func (r UserRepository) UserCreate(ctx context.Context, req types.CreateUserRequest) (u types.User, err error) {
	fmt.Println()
	fmt.Println(req)
	if err := r.db.Create(&req).Error; err != nil {
		return u, err
	}
	r.utils.InjectStructValue(&req, &u)
	return u, nil
}
