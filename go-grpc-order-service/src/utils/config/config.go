package config

import "go-grpc-order-service/src/lib/sql"

type Application struct {
	Config     Config
	SQL        sql.Config
	ServiceUrl ServiceUrl
}

type Config struct {
	Port      string
	JwtSecret string
}

type ServiceUrl struct {
	ProductUrl string
}

func Init() Application {
	return Application{}
}
