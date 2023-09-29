package models

import (
	"errors"
	"strings"
	"time"
)

type Tasks struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status" gorm:"default:false"`
	UserId      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

func (t *Tasks) ValidTask() error {
	if strings.EqualFold(t.Title, "") {
		return errors.New("title is empty")
	}
	return nil
}
