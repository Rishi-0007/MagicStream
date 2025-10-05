package middleware

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/rishi-0007/magicstream-backend/internal/utils"
)

func Auth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth=="" || !strings.HasPrefix(auth,"Bearer "){ c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"missing bearer token"}); return }
		tok := strings.TrimPrefix(auth,"Bearer ")
		claims, err := utils.ParseToken(jwtSecret, tok)
		if err != nil { c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"invalid token"}); return }
		c.Set("userID", claims.UserID); c.Set("role", claims.Role); c.Next()
	}
}
