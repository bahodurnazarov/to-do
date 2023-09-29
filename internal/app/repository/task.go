package repository

import (
	"github.com/bahodurnazarov/to-do/pkg/models"
	"github.com/bahodurnazarov/to-do/pkg/utils"
	"log"
)

func (r *Repository) AddTask(task models.Tasks, authHeader string) error {
	userId, err := utils.FindUserId(authHeader)
	if err != nil {
		log.Println("Can't Find User ID : ", err.Error())
		return err
	}
	task.UserId = userId
	if result := r.Conn.Create(&task); result.Error != nil {
		log.Println("AddTask result : ", result.Error)
		return result.Error
	}
	return nil
}

func (r *Repository) AllTasks(authHeader string) (tasks []models.Tasks, err error) {
	userId, err := utils.FindUserId(authHeader)
	if err != nil {
		log.Println("Can't Find User ID : ", err.Error())
		return nil, err
	}
	if result := r.Conn.Find(&tasks, "user_id = ?", userId); result.Error != nil {
		log.Println("AllTasks result : ", result.Error)
		return tasks, result.Error
	}
	return tasks, nil
}

func (r *Repository) EditTask(updateTask models.Tasks, id string) error {
	var task models.Tasks
	if result := r.Conn.First(&task, id); result.Error != nil {
		log.Println("EditTask result : ", result.Error)
		return result.Error
	}
	task.Status = updateTask.Status
	r.Conn.Save(&task)
	return nil
}

func (r *Repository) RemoveTask(id string) error {
	var task models.Tasks
	if result := r.Conn.First(&task, id); result.Error != nil {
		log.Println("RemoveTask result : ", result.Error)
		return result.Error
	}
	r.Conn.Delete(&task)
	return nil
}
