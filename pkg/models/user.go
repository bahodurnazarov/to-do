package models

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	ID        uint       `gorm:"type:numeric;primary_key"`
	Name      string     `gorm:"type:varchar(100);not null"`
	Email     string     `gorm:"type:varchar(100);uniqueIndex:idx_email;not null"`
	Password  string     `gorm:"type:varchar(100);not null"`
	Role      *string    `gorm:"type:varchar(50);default:'user';not null"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
}

type UsersToken struct {
	ID        int        `gorm:"type:numeric;primary_key"`
	Token     string     `gorm:"type:varchar;not null"`
	UserId    uint       `json:"user_id"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
}

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	Photo           string `json:"photo,omitempty"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Generer")
		return nil, err
	}
	return hashedPassword, nil
}
