package model

import "time"

type Task struct {
	ID        int       `json:"id" gorm:"praimaly_key"`
	Name      string    `json:"name"`
	Context   string    `json:"context"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
