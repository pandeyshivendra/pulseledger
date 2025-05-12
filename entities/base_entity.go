package entities

import (
	"time"
)

type BaseEntity struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
