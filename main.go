package main

import (
	"net"
	"time"

	"github.com/DeniesKresna/jhapi2-user/config"
	conf "github.com/DeniesKresna/jhapi2-user/config"
	"github.com/DeniesKresna/jhapi2-user/gapi"
	"github.com/DeniesKresna/jhapi2-user/pb"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func main() {
	setTimeZone()
	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("unable to load env through env config")
	}
	conf.ProvideConfig()
	conf.ProvideDB()

	config := conf.Get()
	db := conf.GetDB()

	runGrpcServer(&config, db)
}

func runGrpcServer(config *config.ConfigV1, db *gorm.DB) {
	server := gapi.NewServer(config, db)
	grpcServer := grpc.NewServer()
	pb.RegisterApiUserServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", *config.Application.Url)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}

func setTimeZone() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Error().Err(err).Msg("fail to set time zone to Asia/Jakarta")
	}
	time.Local = loc // -> this is setting the global timezone
}
