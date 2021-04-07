package repositories

import (
	"challenge-golang-stone/src/models"
	"database/sql"
	"errors"
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

// SearchByID returns a Account from database based on accountID
func (repository Accounts) SearchByID(accountID uint64) (models.Account, error) {
	rows, err := repository.db.Query(
		"select id, name, cpf, secret, balance, created_at from accounts where id = ?",
		accountID,
	)
	if err != nil {
		return models.Account{}, err
	}
	defer rows.Close()

	var account models.Account

	if rows.Next() {
		if err = rows.Scan(
			&account.ID,
			&account.Name,
			&account.Cpf,
			&account.Secret,
			&account.CreatedAt,
		); err != nil {
			return models.Account{}, err
		}
	} else {
		return models.Account{}, errors.New("Account not found")
	}

	return account, nil
}
