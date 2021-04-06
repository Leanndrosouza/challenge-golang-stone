package validators

import (
	"github.com/Nhanderu/brdoc"
)

// ValidateCPF validate if a cpf is valid
func ValidateCPF(cpf string) bool {
	return brdoc.IsCPF(cpf)
}
