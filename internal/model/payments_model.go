package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Payments struct {
	ID         uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Date       time.Time      `gorm:"not null" json:"date"`
	AmountPaid float64        `gorm:"type:double;not null" json:"amount_paid"`
	Status     string         `gorm:"type:enum('PENDING','SUCCESS','FAILED');not null" json:"status"`
	CreatedAt  time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy  uuid.UUID      `gorm:"type:char(36);not null" json:"created_by"`
	UpdatedAt  time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy  *uuid.UUID     `gorm:"type:char(36)" json:"updated_by"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	DeletedBy  *uuid.UUID     `gorm:"type:char(36)" json:"deleted_by"`
	//
	TransactionID uuid.UUID    `gorm:"type:char(36);not null" json:"transaction_id"`
	Transaction   Transactions `gorm:"foreignKey:TransactionID" json:"transaction"`
}
