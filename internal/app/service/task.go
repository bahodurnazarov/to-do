package service

import (
	"github.com/bahodurnazarov/to-do/pkg/models"
)

func (s *Service) AddTask(task models.Tasks) error {
	if err := task.ValidTask(); err != nil {
		return err
	}
	return s.Repository.AddTask(task)
}

func (s *Service) AllTasks() (task []models.Tasks, err error) {
	tasks, err := s.Repository.AllTasks()
	return tasks, err
}

func (s *Service) EditTask(updateTask models.Tasks, id string) error {
	return s.Repository.EditTask(updateTask, id)
}

func (s *Service) RemoveTask(id string) error {
	return s.Repository.RemoveTask(id)
}
