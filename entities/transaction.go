package entities

import (
	"pulseledger/enums"
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	BaseEntity
	AccountID       uint64              `gorm:"not null;index"`
	OperationTypeID enums.OperationType `gorm:"not null;index"`
	Account         Account             `gorm:"foreignKey:AccountID;constraint:onUpdate:Cascade,OnDelete:Cascade"`
	Amount          decimal.Decimal     `gorm:"type:decimal(20,8);not null"`
	OperationType   OperationType       `gorm:"forignKey:OperationTypeID;constraint:onUpdate:Cascade,OnDelete:Cascade"`
	EventDate       time.Time           `gorm:"not null"`
}
