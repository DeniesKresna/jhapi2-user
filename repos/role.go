package repos

import (
	"context"
	"fmt"

	"github.com/DeniesKresna/jhapi2-user/types"
	"github.com/DeniesKresna/jhapi2/utils"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db    *gorm.DB
	utils utils.IUtils
}

func ProvideRoleRepository(db *gorm.DB, utils utils.IUtils) IRoleRepository {
	return &RoleRepository{
		db:    db,
		utils: utils,
	}
}

func (r RoleRepository) RoleGetByID(ctx context.Context, id int64) (role types.Role, err error) {
	err = r.db.First(&role, id).Error
	return role, err
}

func (r RoleRepository) RoleCreate(ctx context.Context, req types.CreateRoleRequest) (role types.Role, err error) {
	fmt.Println()
	fmt.Println(req)
	if err := r.db.Create(&req).Error; err != nil {
		return role, err
	}
	r.utils.InjectStructValue(&req, &role)
	return role, nil
}
