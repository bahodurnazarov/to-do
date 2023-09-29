package repository

import (
	"fmt"
	"github.com/bahodurnazarov/to-do/pkg/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

// JWTClaims is a struct for storing JWT claims
type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (r *Repository) SignUpUser(newUser models.User) error {
	if result := r.Conn.Create(&newUser); result.Error != nil {
		log.Println("SignUpUser result : ", result.Error)
		return result.Error
	}
	return nil
}

func (r *Repository) SignInUser(payload models.SignInInput) (string, error) {
	var user models.User
	err := r.Conn.First(&user, "email = ?", strings.ToLower(payload.Email)).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("Invalid email or password")
		} else {
			return "", err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return "", fmt.Errorf("Invalid username or password")
	}

	// Generate a JWT token
	expirationTime := time.Now().Add(30 * time.Minute)
	claims := JWTClaims{
		Username: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("dfgdfgw4er5834"))
	if err != nil {
		return "", fmt.Errorf("Failed to generate JWT token")
	}

	var newUser models.UsersToken
	newUser.Token = tokenString
	newUser.UserId = user.ID

	if err := r.Conn.Create(&newUser); err.Error != nil {
		log.Println("SignUpUser result : ", err.Error)
		return "", err.Error
	}

	return tokenString, nil
}
