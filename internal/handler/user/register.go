package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcblastor/api_tweets/internal/dto"
)

func (h *Handler) Register(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.RegisterRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userId, statusCode, err := h.userService.Register(ctx, &req)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, dto.RegisterResponse{
		Id: userId,
	})
}
