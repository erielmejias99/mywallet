package currency

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails/runtime"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type CurrencyI interface {
	GetAll() ([]Currency, error )
	Create(  map[string]interface{} ) ( Currency, error )
	Update(  map[string]interface{} ) error
	Delete(  map[string]interface{} ) error
}

type CurrencyRep struct {
	db * gorm.DB
}

var currencyRep * CurrencyRep = nil

func NewRep(db *gorm.DB) * CurrencyRep {
	if currencyRep == nil{
		currencyRep = &CurrencyRep{ db: db }
	}
	return currencyRep
}

func (cdto * CurrencyRep)GetAll() ([]Currency, error)  {

	var currency []Currency
	result := cdto.db.Find(&currency)
	if result.Error != nil{
		logger := runtime.NewLog()
		logger.New(fmt.Sprintf( "Error getting currency from DB: %s", result.Error.Error() ) )
		return currency, result.Error
	}
	return currency, nil
}

func (cdto * CurrencyRep)Create( map_currency map[string]interface{} ) ( Currency, error ) {

	var currency Currency

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
	result := cdto.db.Create( &currency )
	if result.Error != nil{
		logger := runtime.NewLog()
		logger.New(fmt.Sprintf( "Error getting currency from DB: %s", result.Error.Error() ) )
		return currency, result.Error
	}
	return currency, nil
}

func ( cdto *  CurrencyRep )Update( map_currency map[string]interface{} ) error{

	var currency Currency

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
	result := cdto.db.Create( &currency )
	if result.Error != nil{
		logger := runtime.NewLog()
		logger.New(fmt.Sprintf( "Error getting currency from DB: %s", result.Error.Error() ) )
		return result.Error
	}
	return nil
}

func ( cdto * CurrencyRep )Delete(  map_currency map[string]interface{} ) error {

	var currency Currency

	err := mapstructure.Decode( map_currency, &currency )
	if err != nil{
		return err
	}

	result := cdto.db.Delete( &currency )
	if result.Error != nil{
		logger := runtime.NewLog()
		logger.New(fmt.Sprintf( "Error deleting currency in DB: %s", result.Error.Error() ) )
		return result.Error
	}
	return nil
}