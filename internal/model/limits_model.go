package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Limits struct {
	ID             uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	LimitAvailable float64        `gorm:"type:double;not null" json:"limit_available"`
	CreatedAt      time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy      uuid.UUID      `gorm:"type:char(36);not null" json:"created_by"`
	UpdatedAt      time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy      *uuid.UUID     `gorm:"type:char(36)" json:"updated_by"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	DeletedBy      *uuid.UUID     `gorm:"type:char(36)" json:"deleted_by"`
	//
	ConsumerID uuid.UUID `gorm:"type:char(36);not null" json:"consumer_id"`
	Consumer   Consumers `gorm:"foreignKey:ConsumerID" json:"consumer"`
}
