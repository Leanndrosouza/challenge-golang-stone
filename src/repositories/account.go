package repositories

import (
	"challenge-golang-stone/src/models"
	"database/sql"
)

// Accounts is a repository for accounts
type Accounts struct {
	db *sql.DB
}

// NewAccountRepository returns a new repository of accounts
func NewAccountRepository(db *sql.DB) *Accounts {
	return &Accounts{db}
}

// Create insert a user on database
func (repository Accounts) Create(account models.Account) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into accounts (name, cpf, secret, balance) values (?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(account.Name, account.Cpf, account.Secret, account.Balance)
	if err != nil {
		return 0, err
	}

	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil
}
