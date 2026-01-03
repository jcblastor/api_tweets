package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jcblastor/api_tweets/internal/service/user"
)

type Handler struct {
	api         *gin.Engine
	userService user.UserService
}

func NewHandler(api *gin.Engine, userService user.UserService) *Handler {
	return &Handler{
		api:         api,
		userService: userService,
	}
}

func (h *Handler) RouteList() {
	authRoute := h.api.Group("/auth")
	authRoute.POST("/register", h.Register)
}
