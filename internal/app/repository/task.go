package repository

import (
	"fmt"
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

func (r *Repository) RemoveTask(id string) error {
	var task models.Tasks
	if result := r.Conn.First(&task, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	r.Conn.Delete(&task)
	return nil
}
