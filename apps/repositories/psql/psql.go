package psql

import "gorm.io/gorm"

type psql struct {
	transaction *gorm.DB
}

func NewPsql(transaction *gorm.DB) *psql {
	return &psql{
		transaction: transaction,
	}
}
