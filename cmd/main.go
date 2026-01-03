package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcblastor/api_tweets/internal/config"
	userHandler "github.com/jcblastor/api_tweets/internal/handler/user"
	userRepo "github.com/jcblastor/api_tweets/internal/repository/user"
	userService "github.com/jcblastor/api_tweets/internal/service/user"
	"github.com/jcblastor/api_tweets/pkg/internalsql"
)

func main() {
	r := gin.Default()
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
	userService := userService.NewService(cfg, userRepo)
	userHandler := userHandler.NewHandler(r, userService)
	userHandler.RouteList()

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	r.Run(server)
}
