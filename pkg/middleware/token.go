package middleware

import (
	"github.com/bahodurnazarov/to-do/pkg/db"
	"github.com/bahodurnazarov/to-do/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Check if the Authorization header is present
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			log.Println("Authorization header is missing")
			return
		}

		// Split the Authorization header value
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid bearer token"})
			c.Abort()
			log.Println("Invalid bearer token")
			return
		}

		// Retrieve the token
		token := splitToken[1]
		var usersT models.UsersToken
		conn := db.DB
		result := conn.First(&usersT, "token = ?", token).Error
		if result != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
