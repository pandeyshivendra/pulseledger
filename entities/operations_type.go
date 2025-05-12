package entities

import "time"

type OperationType struct {
	ID           uint8  `gorm:"primaryKey"`
	Description  string `gorm:"not null"`
	IsDepricated bool   `gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
