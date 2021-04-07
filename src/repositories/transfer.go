package repositories

import (
	"challenge-golang-stone/src/models"
	"database/sql"
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
func (repository Transfers) Transfer(transfer models.Transfer) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into transfers (account_origin_id, account_destination_id, amount) values (?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount)
	if err != nil {
		return 0, err
	}

	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil
}
