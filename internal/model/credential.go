package model

import (
	"time"
)

func (Credential) TableName() string {
	return "credential"
}

// nb: if using custom type data, eg=json, array, etc. always define the type for gorm
type Credential struct {
	ID         string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CustomerID string    `gorm:"column:customer_id"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime"`
}
