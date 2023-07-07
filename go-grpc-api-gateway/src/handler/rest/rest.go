package rest

import (
	"context"
	"fmt"
	"go-grpc-api-gateway/docs/swagger"
	"go-grpc-api-gateway/src/client"
	"go-grpc-api-gateway/src/lib/configreader"
	"go-grpc-api-gateway/src/utils/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var once = &sync.Once{}

type REST interface {
	Run()
}

type rest struct {
	http         *gin.Engine
	conf         config.ApplicationMeta
	configreader configreader.Interface
	client       *client.Client
}

func Init(conf config.ApplicationMeta, confReader configreader.Interface, cl *client.Client) REST {
	r := &rest{}
	once.Do(func() {
		httpServ := gin.Default()

		r = &rest{
			conf:         conf,
			configreader: confReader,
			http:         httpServ,
			client:       cl,
		}

		r.http.Use(cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowHeaders:    []string{"*"},
			AllowMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
		}))

		// Set Recovery
		r.http.Use(gin.Recovery())

		r.http.Use(gin.Logger())

		r.Register()
	})

	return r
}

func (r *rest) Run() {
	port := ":8080"

	server := &http.Server{
		Addr:    port,
		Handler: r.http,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(fmt.Sprintf("Serving HTTP error: %s", err.Error()))
		}
	}()
	fmt.Println(fmt.Sprintf("Listening and Serving HTTP on %s", server.Addr))

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	log.Println("Server exiting")
}

func (r *rest) Register() {
	r.registerSwaggerRoutes()
	publicApi := r.http.Group("/public")
	publicApi.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "hello world",
		})
	})

	api := r.http.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "hello world",
		})
	})

	v1.POST("/auth/register", r.RegisterUser)
	v1.POST("/auth/login", r.LoginUser)

	v1.POST("/order", r.AuthRequired, r.CreateOrder)

	v1.POST("/product", r.AuthRequired, r.CreateProduct)
	v1.GET("/product/:id", r.AuthRequired, r.FindOneProduct)
}

func (r *rest) registerSwaggerRoutes() {
	swagger.SwaggerInfo.Title = r.conf.Title
	swagger.SwaggerInfo.Description = r.conf.Description
	swagger.SwaggerInfo.Version = r.conf.Version
	swagger.SwaggerInfo.Host = r.conf.Host
	swagger.SwaggerInfo.BasePath = r.conf.BasePath

	r.http.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
