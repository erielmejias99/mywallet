package models

import "gorm.io/gorm"

type Reason struct {
	gorm.Model
	Label string `gorm:"unique" json:"label" mapstructure:"label" validate:"required,min=5,max=50"`
}
