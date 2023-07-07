package main

import (
	"fmt"
	"go-grpc-auth-service/src/business/domain"
	"go-grpc-auth-service/src/business/usecase"
	authLib "go-grpc-auth-service/src/lib/auth"
	"go-grpc-auth-service/src/lib/configreader"
	"go-grpc-auth-service/src/lib/sql"
	"go-grpc-auth-service/src/pb"
	"go-grpc-auth-service/src/utils/config"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	configFile string = "./etc/cfg/config.json"
)

func main() {
	cfg := config.Init()
	configReader := configreader.Init(configreader.Options{
		ConfigFile: configFile,
	})
	configReader.ReadConfig(&cfg)

	sql := sql.Init(cfg.SQL)

	lis, err := net.Listen("tcp", cfg.Config.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Printf("Auth service running on %s", cfg.Config.Port)

	d := domain.Init(sql)

	al := authLib.Init()

	uc := usecase.Init(al, d)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, uc.Auth)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
