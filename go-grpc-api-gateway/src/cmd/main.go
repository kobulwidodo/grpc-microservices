package main

import (
	"go-grpc-api-gateway/src/client"
	"go-grpc-api-gateway/src/handler/rest"
	"go-grpc-api-gateway/src/lib/configreader"
	"go-grpc-api-gateway/src/utils/config"
)

// @contact.name   Rakhmad Giffari Nurfadhilah
// @contact.url    https://fadhilmail.tech/
// @contact.email  rakhmadgiffari14@gmail.com

// @securitydefinitions.apikey BearerAuth
// @in header
// @name Authorization

const (
	configFile string = "./etc/cfg/config.json"
)

func main() {
	cfg := config.Init()
	configReader := configreader.Init(configreader.Options{
		ConfigFile: configFile,
	})
	configReader.ReadConfig(&cfg)

	client := client.Init(cfg.ServiceUrl)

	r := rest.Init(cfg.Meta, configReader, client)

	r.Run()
}
