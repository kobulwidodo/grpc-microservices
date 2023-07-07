package config

import (
	"time"
)

type Application struct {
	Meta       ApplicationMeta
	Gin        GinConfig
	ServiceUrl ServiceUrl
}

type ApplicationMeta struct {
	Title       string
	Description string
	Host        string
	BasePath    string
	Version     string
}

type GinConfig struct {
	Port            string
	Mode            string
	Timeout         time.Duration
	ShutdownTimeout time.Duration
	CORS            CORSConfig
	Meta            ApplicationMeta
}

type ServiceUrl struct {
	AuthServiceUrl    string
	ProductServiceUrl string
	OrderServiceUrl   string
}

type CORSConfig struct {
	Mode string
}

func Init() Application {
	return Application{}
}
