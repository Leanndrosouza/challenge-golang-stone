package models

import (
	"errors"
	"time"
)

// Transfer model for transfers
type Transfer struct {
	ID                   uint64    `json:"id,omitempty"`
	AccountOriginID      uint64    `json:"account_origin_id,omitempty"`
	AccountDestinationID uint64    `json:"account_destination_id,omitempty"`
	Amount               int       `json:"amount,omitempty"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
}

// Prepare validate transfer fields
func (transfer *Transfer) Prepare() error {
	if err := transfer.validate(); err != nil {
		return err
	}

	return nil
}

func (transfer *Transfer) validate() error {
	if transfer.AccountOriginID == 0 {
		return errors.New("AccountOriginID is a required field")
	}
	if transfer.AccountDestinationID == 0 {
		return errors.New("AccountDestinationID is a required field")
	}
	if transfer.Amount == 0 {
		return errors.New("Transfer Amount must be non-zero")
	}

	return nil
}
