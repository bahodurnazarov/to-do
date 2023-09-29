package service

import (
	"fmt"
	"github.com/bahodurnazarov/to-do/pkg/models"
	"strings"
)

func (s *Service) SignUpUser(payload models.SignUpInput) error {
	if payload.Password != payload.PasswordConfirm {
		return fmt.Errorf("password is not mutch")
	}

	hashedPassword, err := models.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
	}

	return s.Repository.SignUpUser(newUser)
}

func (s *Service) SignInUser(payload models.SignInInput) (string, error) {
	token, err := s.Repository.SignInUser(payload)
	if err != nil {
		return "", err
	}
	return token, nil
}
