package main

import (
	"fmt"
	"go-grpc-product-service/src/business/domain"
	"go-grpc-product-service/src/business/usecase"
	"go-grpc-product-service/src/lib/configreader"
	"go-grpc-product-service/src/lib/sql"
	"go-grpc-product-service/src/pb"
	"go-grpc-product-service/src/utils/config"
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

	uc := usecase.Init(d)

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, uc.Product)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
