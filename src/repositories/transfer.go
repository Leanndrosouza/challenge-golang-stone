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
