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
func (account *Account) Prepare(step string) error {
	if err := account.validate(step); err != nil {
		return err
	}
	if err := account.format(step); err != nil {
		return err
	}

	return nil
}

func (account *Account) validate(step string) error {
	if step == "create" && strings.TrimSpace(account.Name) == "" {
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

func (account *Account) format(step string) error {
	account.Name = strings.TrimSpace(account.Name)
	account.Cpf = strings.TrimSpace(account.Cpf)

	reg, err := regexp.Compile("\\D")
	if err != nil {
		return err
	}

	account.Cpf = reg.ReplaceAllString(account.Cpf, "")

	if step == "create" {
		hashedSecret, err := security.Hash(account.Secret)
		if err != nil {
			return err
		}

		account.Secret = string(hashedSecret)
		account.Balance = 0
	}

	return nil
}
