package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Phone     string
	ImageUrl  string
	Password  string
	RoleId    uint64
}

type CreateUserRequest struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
	Phone     string `validate:"numeric,min=11,max=15"`
	ImageUrl  string `validate:"-"`
	Password  string `validate:"required,min=8"`
	RoleId    uint64 `validate:"required,numeric"`
}

func (u *CreateUserRequest) TableName() string {
	return "users"
}
