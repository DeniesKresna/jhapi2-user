package gapi

import (
	"github.com/DeniesKresna/jhapi2-user/config"
	"github.com/DeniesKresna/jhapi2-user/handlers"
	"github.com/DeniesKresna/jhapi2-user/pb"
	"github.com/DeniesKresna/jhapi2-user/repos"
	"github.com/DeniesKresna/jhapi2/utils"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedApiUserServer
	config      *config.ConfigV1
	v           *validator.Validate
	userHandler handlers.IUserHandler
	roleHandler handlers.IRoleHandler
	utils       utils.IUtils
}

func NewServer(config *config.ConfigV1, db *gorm.DB) *Server {
	utils := utils.Provide()

	//inject list repo and handler to server
	userRepo := repos.ProvideUserRepository(db, utils)
	userHandler := handlers.ProvideUserHandler(userRepo)
	roleRepo := repos.ProvideRoleRepository(db, utils)
	roleHandler := handlers.ProvideRoleHandler(roleRepo)

	server := &Server{
		config:      config,
		v:           validator.New(),
		userHandler: userHandler,
		roleHandler: roleHandler,
		utils:       utils,
	}

	return server
}
