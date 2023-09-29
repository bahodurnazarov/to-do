package utils

import (
	"fmt"
	"github.com/bahodurnazarov/to-do/pkg/db"
	"github.com/bahodurnazarov/to-do/pkg/models"
	"strings"
)

func FindUserId(authHeader string) (uint, error) {
	if authHeader == "" {
		fmt.Errorf("authHeader is empty")
		return 0, nil
	}

	// Split the Authorization header value
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		fmt.Errorf("Invalid To Split Token")
		return 0, nil
	}

	var usersT models.UsersToken
	// Retrieve the token
	token := splitToken[1]
	conn := db.DB
	result := conn.First(&usersT, "token = ?", token).Error
	if result != nil {
		fmt.Errorf("Invalid token")
		return 0, nil
	}
	return usersT.UserId, nil
}
