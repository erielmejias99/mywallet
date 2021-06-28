package transaction

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails/runtime"
	"gorm.io/gorm"
)

type TransactionI interface {
	Create( map[string]interface{} ) (Transaction, error)
}

type TransactionRep struct {
	db * gorm.DB
}

var transactionRep * TransactionRep = nil
func GetTransactionRep( db * gorm.DB) *TransactionRep{
	if transactionRep == nil{
		transactionRep = &TransactionRep{ db: db }
	}
	return transactionRep
}

func (receiver TransactionRep) Create(transaction map[string]interface{}) (Transaction, error) {
	var newTransaction Transaction

	// Decode javaScript object
	err := mapstructure.Decode( transaction, &newTransaction )
	if err != nil{
		return newTransaction, err
	}

	// Validate fields
	val := validator.New()
	err = val.Struct( &newTransaction )
	if err != nil{
		return newTransaction, err
	}

	// Save Model in DataBase
	result := receiver.db.Create( &newTransaction )
	if result.Error != nil {
		logger := runtime.NewLog()
		logger.New(fmt.Sprintf( "Error getting currency from DB: %s", result.Error.Error() ) )
		return newTransaction, result.Error
	}

	return newTransaction, nil

}