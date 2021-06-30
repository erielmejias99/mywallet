package controllers

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails/runtime"
	"gorm.io/gorm"
	"wallet/models"
)

type CurrencyController struct{
	Db * gorm.DB
}

func (cdto *CurrencyController)GetAll() ([]models.Currency, error)  {

	var currency []models.Currency
	result := cdto.Db.Find(&currency)
	if result.Error != nil{
		logger := runtime.NewLog()
		logger.New(fmt.Sprintf( "Error getting currency from DB: %s", result.Error.Error() ) )
		return currency, result.Error
	}
	return currency, nil
}

func (cdto *CurrencyController)Create( map_currency map[string]interface{} ) (models.Currency, error ) {

	var currency models.Currency

	err := mapstructure.Decode( map_currency, &currency )
	if err != nil{
		return currency,err
	}

	val := validator.New()

	err = val.Struct( currency )
	if err != nil{
		return currency, err
	}

	//var currency []Currency
	result := cdto.Db.Create( &currency )
	if result.Error != nil{
		logger := runtime.NewLog()
		logger.New(fmt.Sprintf( "Error getting currency from DB: %s", result.Error.Error() ) )
		return currency, result.Error
	}
	return currency, nil
}

func ( cdto *CurrencyController)Update( map_currency map[string]interface{} ) error{

	var currency models.Currency

	err := mapstructure.Decode( map_currency, &currency )
	if err != nil{
		return err
	}

	val := validator.New()

	err = val.Struct( currency )
	if err != nil{
		return err
	}

	//var currency []Currency
	result := cdto.Db.Create( &currency )
	if result.Error != nil{
		logger := runtime.NewLog()
		logger.New(fmt.Sprintf( "Error getting currency from DB: %s", result.Error.Error() ) )
		return result.Error
	}
	return nil
}

func ( cdto *CurrencyController)Delete(  map_currency map[string]interface{} ) error {

	var currency models.Currency

	err := mapstructure.Decode( map_currency, &currency )
	if err != nil{
		return err
	}

	result := cdto.Db.Delete( &currency )
	if result.Error != nil{
		logger := runtime.NewLog()
		logger.New(fmt.Sprintf( "Error deleting currency in DB: %s", result.Error.Error() ) )
		return result.Error
	}
	return nil
}