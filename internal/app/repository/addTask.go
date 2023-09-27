package repository

import (
	"fmt"
	"github.com/bahodurnazarov/to-do/pkg/models"
)

func (r *Repository) AddTask(task models.Tasks) error {
	if result := r.Conn.Create(&task); result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}
