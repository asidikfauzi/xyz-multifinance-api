package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Consumers struct {
	ID           uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	NIK          string         `gorm:"type:varchar(45);unique;not null" json:"nik"`
	FullName     string         `gorm:"type:varchar(45);not null" json:"full_name"`
	LegalName    string         `gorm:"type:varchar(45);not null" json:"legal_name"`
	Phone        string         `gorm:"type:varchar(45);not null" json:"phone"`
	PlaceOfBirth string         `gorm:"type:varchar(45);not null" json:"place_of_birth"`
	DateOfBirth  string         `gorm:"type:varchar(45);not null" json:"date_of_birth"`
	Salary       string         `gorm:"type:varchar(45);not null" json:"salary"`
	KTPImage     string         `gorm:"type:varchar(255)" json:"ktp_image"`
	SelfieImage  string         `gorm:"type:varchar(255)" json:"selfie_image"`
	CreatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	//
	UserID      uuid.UUID      `gorm:"type:char(36);not null" json:"user_id"`
	User        Users          `gorm:"foreignKey:UserID" json:"user"`
	Limits      []Limits       `gorm:"foreignKey:ConsumerID" json:"limits"`
	Transaction []Transactions `gorm:"foreignKey:ConsumerID" json:"transactions"`
}
