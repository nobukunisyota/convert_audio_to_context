package model

import "time"

type Task struct {
	ID        int       `json:"id" gorm:"praimaly_key"`
	Name      string    `json:"name"`
	Context   string    `json:"context"`
	Tag       string    `json:"tag"`
	UpdatedAt time.Time `json:"updated_at"`
}
