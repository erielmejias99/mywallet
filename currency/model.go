package currency

import (
	"gorm.io/gorm"
)

type Currency struct{
	Name string `gorm:"unique" json:"name" mapstructure:"name" validate:"required,min=3,max=5"`
	Balance uint64 `gorm:"not null" json:"balance" mapstructure:"balance" validate:"min=20"`
	USDChange float32 `gorm:"not null" json:"usd_change" mapstructure:"usd_change" validate:"-"`
	gorm.Model
}

func (c * Currency )eriel() int {
	return 15
}

