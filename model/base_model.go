package model

import "time"

type BaseModel struct {
	ID        uint       `json:"id" gorm:"primary_key;not null"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
