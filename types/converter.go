package types

import (
	pbRole "github.com/DeniesKresna/jhapi2-user/pb/role"
	pbUser "github.com/DeniesKresna/jhapi2-user/pb/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertUser(user User) *pbUser.User {
	return &pbUser.User{
		Id:        uint64(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		ImageUrl:  user.ImageUrl,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func ConvertRole(role Role) *pbRole.Role {
	return &pbRole.Role{
		Id:        uint64(role.ID),
		Name:      role.Name,
		CreatedAt: timestamppb.New(role.CreatedAt),
		UpdatedAt: timestamppb.New(role.UpdatedAt),
	}
}
