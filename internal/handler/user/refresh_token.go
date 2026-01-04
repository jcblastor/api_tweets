package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcblastor/api_tweets/internal/dto"
)

func (h *Handler) RefreshToken(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.RefreshTokenRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userId := c.GetInt64("userId")

	token, refreshToken, statusCode, err := h.userService.RefreshToken(ctx, &req, userId)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(statusCode, dto.RefreshTokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
}
