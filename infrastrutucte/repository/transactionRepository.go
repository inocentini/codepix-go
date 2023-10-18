package repository

import (
	"fmt"

	"github.com/inocentini/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (transactionRepo *TransactionRepositoryDb) RegisterTransaction(transaction *model.Transaction) error {
	err := transactionRepo.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (transactionRepo *TransactionRepositoryDb) SaveTransaction(transaction *model.Transaction) error {
	err := transactionRepo.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (transactionRepo *TransactionRepositoryDb) FindTransaction(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	transactionRepo.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}
	return &transaction, nil
}
