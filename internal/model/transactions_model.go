package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transactions struct {
	ID             uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	ContractNumber string         `gorm:"type:varchar(45);unique;not null" json:"contract_number"`
	OTR            float64        `gorm:"type:double;not null" json:"otr"`
	Tenor          int            `gorm:"type:int;not null" json:"tenor"`
	AdminFee       float64        `gorm:"type:double;not null" json:"admin_fee"`
	InstallmentAmt float64        `gorm:"type:double;not null" json:"installment_amt"`
	AmountInterest float64        `gorm:"type:double;not null" json:"amount_interest"`
	AssetName      string         `gorm:"type:varchar(45);not null" json:"asset_name"`
	CreatedAt      time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy      uuid.UUID      `gorm:"type:char(36);not null" json:"created_by"`
	UpdatedAt      time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy      *uuid.UUID     `gorm:"type:char(36)" json:"updated_by"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	DeletedBy      *uuid.UUID     `gorm:"type:char(36)" json:"deleted_by"`
	//
	ConsumerID uuid.UUID  `gorm:"type:char(36);not null" json:"consumer_id"`
	Consumer   Consumers  `gorm:"foreignKey:ConsumerID" json:"consumer"`
	Payments   []Payments `gorm:"foreignKey:TransactionID" json:"payments"`
}
