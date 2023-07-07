package main

import (
	"fmt"
	"go-grpc-order-service/src/business/domain"
	"go-grpc-order-service/src/business/usecase"
	"go-grpc-order-service/src/client"
	"go-grpc-order-service/src/lib/configreader"
	"go-grpc-order-service/src/lib/sql"
	"go-grpc-order-service/src/pb"
	"go-grpc-order-service/src/utils/config"
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

	client := client.Init(cfg.ServiceUrl)

	uc := usecase.Init(d, client)

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, uc.Order)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
