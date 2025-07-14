package repository

import "e-money-service/internal/model"

type AccountRepository interface {
	CreateAccount(account *model.Account) error
	CreateTransaction(accountID string, transaction *model.Transaction) error
	FindAccountByUsername(username string) (*model.Transaction, error)
	UpdateAccount(account *model.Account) error
}

type AccountRepositoryImpl struct {
	// Add any necessary fields, such as a database connection or context

}
