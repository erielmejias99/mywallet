package main

import (
	"gorm.io/gorm"
	"wallet/models"
)

func InitMigrations(db * gorm.DB) error{
	var err error
	err = db.AutoMigrate( &models.Currency{} )
	if err != nil{
		return err
	}
	err = db.AutoMigrate( &models.Transaction{} )
	if err != nil{
		return err
	}

	return nil
}
