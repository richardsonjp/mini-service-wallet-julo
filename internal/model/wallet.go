package model

import (
	"go-skeleton/internal/model/enum"
	"time"
)

func (Wallet) TableName() string {
	return "wallet"
}

type WalletStatus = enum.WalletStatus

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Wallet struct {
	ID         string       `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CustomerID string       `gorm:"column:customer_id"`
	Balance    int64        `gorm:"column:balance"`
	Status     WalletStatus `gorm:"column:status"`
	CreatedAt  time.Time    `gorm:"column:created_at;type:datetime"`
	UpdatedAt  time.Time    `gorm:"column:updated_at;type:datetime"`
}
