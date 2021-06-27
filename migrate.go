package main

import (
	"gorm.io/gorm"
	"wallet/currency"
	"wallet/transaction"
)

func InitMigrations(db * gorm.DB) error{
	var err error
	err = db.AutoMigrate( &currency.Currency{} )
	if err != nil{
		return err
	}
	err = db.AutoMigrate( &transaction.Reason{} )
	if err != nil{
		return err
	}
	err = db.AutoMigrate( &transaction.Transaction{} )
	if err != nil{
		return err
	}

	return nil
}
