package config

import "go-grpc-product-service/src/lib/sql"

type Application struct {
	Config Config
	SQL    sql.Config
}

type Config struct {
	Port      string
	JwtSecret string
}

type CORSConfig struct {
	Mode string
}

func Init() Application {
	return Application{}
}
