package currency

import (
	"gorm.io/gorm"
)

type Currency struct{
	Name string `gorm:"unique" json:"name"`
	Balance uint64 `gorm:"not null" json:"balance"`
	USDChange float32 `gorm:"not null" json:"usd_change"`
	gorm.Model
}


