package service

import (
	"github.com/bahodurnazarov/to-do/pkg/models"
)

func (s *Service) AddTask(task models.Tasks) error {
	if err := task.Valid(); err != nil {
		return err
	}
	return s.Repository.AddTask(task)
}
