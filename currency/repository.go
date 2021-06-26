package currency

import (
	"fmt"
	"github.com/wailsapp/wails/runtime"
	"gorm.io/gorm"
)

type CurrencyI interface {
	GetAll() ([]Currency, error )
	Create( currency * Currency ) error
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

func (cdto * CurrencyRep)Create( currency * Currency) error  {

	//var currency []Currency
	result := cdto.db.Create( currency )
	if result.Error != nil{
		logger := runtime.NewLog()
		logger.New(fmt.Sprintf( "Error getting currency from DB: %s", result.Error.Error() ) )
		return result.Error
	}
	return nil
}