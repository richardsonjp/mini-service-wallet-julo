package model

import (
	"go-skeleton/internal/model/enum"
	"time"
)

func (Transaction) TableName() string {
	return "transaction"
}

type TransactionType = enum.TransactionType
type GenericStatus = enum.GenericStatus

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Transaction struct {
	ID          string          `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	WalletID    string          `gorm:"column:wallet_id"`
	Amount      int64           `gorm:"column:amount"`
	Status      GenericStatus   `gorm:"column:status"`
	Type        TransactionType `gorm:"column:type"`
	ReferenceID string          `gorm:"column:reference_id"`
	CreatedAt   time.Time       `gorm:"column:created_at;type:datetime"`
}
