package repository

import (
	"github.com/bahodurnazarov/to-do/pkg/models"
	"log"
)

func (r *Repository) SignUpUser(newUser models.User) error {
	if result := r.Conn.Create(&newUser); result.Error != nil {
		log.Println("SignUpUser result : ", result.Error)
		return result.Error
	}
	return nil
}
