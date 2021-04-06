package models

import "time"

// Account model for accounts
type Account struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Cpf       string    `json:"cpf,omitempty"`
	Secret    string    `json:"secret,omitempty"`
	Balance   int       `json:"balance,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
