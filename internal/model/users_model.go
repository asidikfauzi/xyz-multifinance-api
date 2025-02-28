package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Email     string         `gorm:"type:varchar(45);unique;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	//
	RoleID   uuid.UUID   `gorm:"type:char(36);not null" json:"role_id"`
	Role     Roles       `gorm:"foreignKey:RoleID" json:"role"`
	Consumer []Consumers `gorm:"foreignKey:UserID" json:"consumers"`
}
