package transaction

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails/runtime"
	"gorm.io/gorm"
	"math"
	"wallet/currency"
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

	var currencyWallet currency.Currency
	resp := receiver.db.Find( &currencyWallet, newTransaction.CurrencyID )
	if resp.Error != nil{
		return newTransaction, resp.Error
	}

	// Check if the amount of the transaction is negative and there is such amount.
	if newTransaction.Amount < 0{
		if currencyWallet.Balance >= int(math.Abs( float64( newTransaction.Amount ) )) {
			return newTransaction, errors.New("not enough balance available")
		}
	}

	// Update the currency balance
	currencyWallet.Balance += newTransaction.Amount

	//validate after update field
	err = val.Struct(&currencyWallet)
	if err != nil {
		return newTransaction,err
	}

	resp = reasonRepository.db.Update( "balance", currencyWallet )
	if resp.Error != nil {
		return newTransaction,resp.Error
	}

	return newTransaction, nil

}