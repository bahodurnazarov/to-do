package models

import (
	"errors"
	"time"
)

type Tasks struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status" gorm:"default:false"`
	CreatedAt   time.Time `json:"autoCreateAt"`
}

func (t *Tasks) Valid() error {
	if t.Title == "" {
		return errors.New("gfgfgd")
	}
	return nil
}
