package dto

import (
	"pulseledger/enums"

	"github.com/shopspring/decimal"
)

type CreateTransactionRequest struct {
	AccountID       uint64              `json:"account_id" validate:"required,gt=0"`
	OperationTypeID enums.OperationType `json:"operation_type_id" validate:"required"`
	Amount          decimal.Decimal     `json:"amount" validate:"required"`
}
