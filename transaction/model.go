package transaction

import (
	"gorm.io/gorm"
)

type Reason struct {
	gorm.Model
	Label string `gorm:"unique" json:"label" mapstructure:"label" validate:"required,min=5,max=50"`
	TransactionID uint64 `json:"transaction_id" mapstructure:"transaction_id" validate:"-"`
}

type Transaction struct{
	gorm.Model
	CurrencyID uint64 `json:"currency_id" mapstructure:"currency_id" validate:"-"`
	Amount int `gorm:"" json:"amount" mapstructure:"amount" validate:"required,min=0"`
	Description string `gorm:"" json:"description" mapstructure:"description" validate:"max=255"`
	Reason Reason
}
