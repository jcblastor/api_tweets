package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jcblastor/api_tweets/internal/config"
	postHandler "github.com/jcblastor/api_tweets/internal/handler/post"
	userHandler "github.com/jcblastor/api_tweets/internal/handler/user"
	postRepo "github.com/jcblastor/api_tweets/internal/repository/post"
	userRepo "github.com/jcblastor/api_tweets/internal/repository/user"
	postService "github.com/jcblastor/api_tweets/internal/service/post"
	userService "github.com/jcblastor/api_tweets/internal/service/user"
	"github.com/jcblastor/api_tweets/pkg/internalsql"
)

func main() {
	r := gin.Default()
	validate := validator.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := internalsql.ConnectMySQL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/check-health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "it's works",
		})
	})

	userRepo := userRepo.NewRepository(db)
	postRepo := postRepo.NewPostRepository(db)

	userService := userService.NewService(cfg, userRepo)
	postService := postService.NewPostService(cfg, postRepo)

	userHandler := userHandler.NewHandler(r, validate, userService)
	postHandler := postHandler.NewHandler(r, validate, postService)

	userHandler.RouteList(cfg.SecretJwt)
	postHandler.RouterList(cfg.SecretJwt)

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
