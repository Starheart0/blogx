package models

import (
	"time"
)

type Model struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time
}

type IDRequest struct {
	ID uint `json:"ID" form:"id" uri:"id"`
}

type RemoveRequest struct {
	IDlist []uint `json:"IDlist"`
}
