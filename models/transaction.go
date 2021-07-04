package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CurrencyID  uint   `json:"currency_id" mapstructure:"currency_id" validate:"-"`
	Amount      int    `gorm:"" json:"amount" mapstructure:"amount" validate:"required"`
	Description string `gorm:"" json:"description" mapstructure:"description" validate:"max=255"`
	Reason      string `gorm:"" json:"reason" mapstructure:"reason" validate:"required,min=3,max=50"`
}