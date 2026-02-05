package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               uint           `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Country          string         `gorm:"column:country;size:255;not null" json:"country"`
	CreditCardType   string         `gorm:"column:credit_card_type;size:255;not null" json:"credit_card_type"`
	CreditCardNumber string         `gorm:"column:credit_card_number;size:255;not null" json:"credit_card"`
	FirstName        string         `gorm:"column:first_name;size:255;not null" json:"first_name"`
	LastName         string         `gorm:"column:last_name;size:255;not null" json:"last_name"`
	CreatedAt        time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

type Spending struct {
	ID        uint           `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID    uint           `gorm:"column:user_id;size:255;not null" json:"user_id"`
	TotalBuy  int64          `gorm:"column:total_buy;size:255;not null" json:"total_buy"`
	User      User           `json:"user"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}
