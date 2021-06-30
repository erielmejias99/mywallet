package currency

import (
	"gorm.io/gorm"
	"wallet/transaction"
)

type Currency struct{
	Name string `gorm:"unique" json:"name" mapstructure:"name" validate:"required,min=3,max=5"`
	Balance int `gorm:"not null" json:"balance" mapstructure:"balance" validate:"eq=0"`
	USDChange float32 `gorm:"not null" json:"usd_change" mapstructure:"usd_change" validate:"min=0"`
	Transactions []transaction.Transaction
	gorm.Model
}