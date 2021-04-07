package repositories

import (
	"challenge-golang-stone/src/models"
	"database/sql"
	"errors"
)

// Transfers is a repository for transfers
type Transfers struct {
	db *sql.DB
}

// NewTransferRepository returns a new repository of transfers
func NewTransferRepository(db *sql.DB) *Transfers {
	return &Transfers{db}
}

// GetByOriginID return all transfers from database based account_origin_id
func (repository Transfers) GetByOriginID(originID uint64) ([]models.Transfer, error) {
	rows, err := repository.db.Query(
		"select id, account_origin_id, account_destination_id, amount, created_at from transfers where account_origin_id = ?",
		originID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transfers []models.Transfer

	for rows.Next() {
		var transfer models.Transfer

		if err = rows.Scan(
			&transfer.ID,
			&transfer.AccountOriginID,
			&transfer.AccountDestinationID,
			&transfer.Amount,
			&transfer.CreatedAt,
		); err != nil {
			return nil, err
		}

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

// Transfer make a transfer between two accounts
func (repository Transfers) Transfer(transfer models.Transfer) error {
	tx, err := repository.db.Begin()
	if err != nil {
		return err
	}

	newOriginBalance, newDestinationBalance, err := repository.amountAfterTransfer(transfer)
	if err != nil {
		return err
	}

	{
		stmt, err := tx.Prepare("insert into transfers (account_origin_id, account_destination_id, amount) values (?,?,?)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount); err != nil {
			tx.Rollback()
			return err
		}
	}

	{
		stmt, err := tx.Prepare("update accounts set balance = ? where id = ?")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(newOriginBalance, transfer.AccountOriginID); err != nil {
			tx.Rollback()
			return err
		}
	}

	{
		stmt, err := tx.Prepare("update accounts set balance = ? where id = ?")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(newDestinationBalance, transfer.AccountDestinationID); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// CanTransfer returns a error if transfer can't be done
func (repository Transfers) CanTransfer(transfer models.Transfer) error {
	if transfer.AccountDestinationID == transfer.AccountOriginID {
		return errors.New("Invalid destination account")
	}

	if transfer.Amount <= 0 {
		return errors.New("Invalid transfer amount")
	}

	accountRepository := NewAccountRepository(repository.db)

	originAccount, err := accountRepository.SearchByID(transfer.AccountOriginID)
	if err != nil {
		if err.Error() == "Account not found" {
			return errors.New("Origin Account not found")
		}
		return err
	}

	_, err = accountRepository.SearchByID(transfer.AccountDestinationID)
	if err != nil {
		if err.Error() == "Account not found" {
			return errors.New("Destination Account not found")
		}
		return err
	}

	if originAccount.Balance < transfer.Amount {
		return errors.New("No balance to make this transfer")
	}

	return nil
}

func (repository Transfers) amountAfterTransfer(transfer models.Transfer) (int, int, error) {
	accountRepository := NewAccountRepository(repository.db)

	originAccount, err := accountRepository.SearchByID(transfer.AccountOriginID)
	if err != nil {
		if err.Error() == "Account not found" {
			return 0, 0, errors.New("Origin Account not found")
		}
		return 0, 0, err
	}

	detinationAccount, err := accountRepository.SearchByID(transfer.AccountDestinationID)
	if err != nil {
		if err.Error() == "Account not found" {
			return 0, 0, errors.New("Destination Account not found")
		}
		return 0, 0, err
	}

	return originAccount.Balance - transfer.Amount,
		detinationAccount.Balance + transfer.Amount,
		nil
}
