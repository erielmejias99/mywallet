package transaction

import "gorm.io/gorm"

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
	return Transaction{}, nil
}