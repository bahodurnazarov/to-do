package middleware

import (
	"github.com/bahodurnazarov/to-do/pkg/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var cfg *models.Settings
		cfg := config.Cnfg
		secretKey := cfg.App.Token
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

		log.Println("111111111111111111", splitToken)
		// Retrieve the token
		token := splitToken[1]
		log.Println("TOKEN00000000", token)
		log.Println("secretKey2222222222", secretKey)
		// Validate the token (you can implement your own validation logic here)
		if token != secretKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			log.Println("Invalid token")
			return
		}

		c.Next()
	}
}
