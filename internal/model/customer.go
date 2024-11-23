package model

import (
	"time"
)

func (Customer) TableName() string {
	return "customer"
}

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Customer struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	XID       string    `gorm:"column:xid"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
}
