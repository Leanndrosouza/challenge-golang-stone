package models

import (
	"challenge-golang-stone/src/validators"
	"errors"
	"regexp"
	"strings"
)

// Login model for accounts
type Login struct {
	Cpf    string `json:"cpf,omitempty"`
	Secret string `json:"secret,omitempty"`
}

// Prepare format and validate Login fields
func (login *Login) Prepare() error {
	if err := login.validate(); err != nil {
		return err
	}
	if err := login.format(); err != nil {
		return err
	}

	return nil
}

func (login *Login) validate() error {
	if strings.TrimSpace(login.Cpf) == "" {
		return errors.New("CPF is a required field")
	}
	if !validators.ValidateCPF(login.Cpf) {
		return errors.New("CPF invalid")
	}
	if strings.TrimSpace(login.Secret) == "" {
		return errors.New("Secret is a required field")
	}

	return nil
}

func (login *Login) format() error {
	login.Cpf = strings.TrimSpace(login.Cpf)

	reg, err := regexp.Compile("\\D")
	if err != nil {
		return err
	}

	login.Cpf = reg.ReplaceAllString(login.Cpf, "")

	return nil
}
