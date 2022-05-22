package types

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string
}

type CreateRoleRequest struct {
	Name string `validate:"required"`
}

func (u *CreateRoleRequest) TableName() string {
	return "roles"
}
