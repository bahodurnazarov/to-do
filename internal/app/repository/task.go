package repository

import (
	"github.com/bahodurnazarov/to-do/pkg/models"
	"log"
)

func (r *Repository) AddTask(task models.Tasks) error {
	if result := r.Conn.Create(&task); result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (r *Repository) AllTasks() (tasks []models.Tasks, err error) {
	if result := r.Conn.Find(&tasks); result.Error != nil {
		log.Println(result.Error)
		return tasks, result.Error
	}
	return tasks, nil
}
