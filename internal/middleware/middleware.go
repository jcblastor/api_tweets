package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jcblastor/api_tweets/pkg/jwt"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userId, username, err := jwt.ValidateToken(header, secretKey, true)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		c.Set("userId", userId)
		c.Set("username", username)
		c.Next()
	}
}

func AuthRefreshTokenMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userId, username, err := jwt.ValidateToken(header, secretKey, false)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		c.Set("userId", userId)
		c.Set("username", username)
		c.Next()
	}
}
