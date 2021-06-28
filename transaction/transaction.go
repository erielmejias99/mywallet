package transaction

import (
	"errors"
	"gorm.io/gorm"
	"math"
	currency "wallet/currency"
)

type Transaction struct{
	gorm.Model
	CurrencyID uint `json:"currency_id" mapstructure:"currency_id" validate:"-"`
	Amount int `gorm:"" json:"amount" mapstructure:"amount" validate:"required,min=0"`
	Description string `gorm:"" json:"description" mapstructure:"description" validate:"max=255"`
	Reason Reason
}

func ( t *Transaction ) GetCurrency(tx*gorm.DB) (*currency.Currency, error )  {
	var currencyWallet currency.Currency
	resp := tx.Find( &currencyWallet, t.CurrencyID )
	if resp.Error != nil{
		return &currencyWallet, resp.Error
	}
	return &currencyWallet, nil
}

func (t *Transaction) AfterSave(tx*gorm.DB) ( error )  {

	currencyWallet, err := t.GetCurrency(tx)
	if err != nil{
		return errors.New("error getting Currency of the transaction")
	}

	// Check if the amount of the transaction is negative and there is such amount.
	if t.Amount < 0{
		if currencyWallet.Balance >= int(math.Abs( float64( t.Amount ) )) {
			return errors.New("not enough balance available")
		}
	}

	// Update the currency balance
	currencyWallet.Balance += t.Amount

	resp := tx.Update( "balance", currencyWallet )
	if resp.Error != nil {
		return resp.Error
	}

	return nil
}
