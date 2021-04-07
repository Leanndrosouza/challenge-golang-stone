package models

import (
	"challenge-golang-stone/src/security"
	"challenge-golang-stone/src/validators"
	"errors"
	"regexp"
	"strings"
	"time"
)

// Account model for accounts
type Account struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Cpf       string    `json:"cpf,omitempty"`
	Secret    string    `json:"secret,omitempty"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare format and validate account fields
func (account *Account) Prepare() error {
	if err := account.validate(); err != nil {
		return err
	}
	if err := account.format(); err != nil {
		return err
	}

	return nil
}

func (account *Account) validate() error {
	if strings.TrimSpace(account.Name) == "" {
		return errors.New("Name is a required field")
	}
	if strings.TrimSpace(account.Cpf) == "" {
		return errors.New("CPF is a required field")
	}
	if !validators.ValidateCPF(account.Cpf) {
		return errors.New("CPF invalid")
	}
	if strings.TrimSpace(account.Secret) == "" {
		return errors.New("Secret is a required field")
	}

	return nil
}

func (account *Account) format() error {
	account.Name = strings.TrimSpace(account.Name)
	account.Cpf = strings.TrimSpace(account.Cpf)

	reg, err := regexp.Compile("\\D")
	if err != nil {
		return err
	}

	account.Cpf = reg.ReplaceAllString(account.Cpf, "")

	hashedSecret, err := security.Hash(account.Secret)
	if err != nil {
		return err
	}

	account.Secret = string(hashedSecret)
	account.Balance = 0

	return nil
}
