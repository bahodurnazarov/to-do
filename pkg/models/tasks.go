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
	CreatedAt   time.Time `json:"CreatedAt"`
}

func (t *Tasks) Valid() error {
	if strings.EqualFold(t.Title, "") {
		return errors.New("title is empty")
	}
	return nil
}
