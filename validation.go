// Package validation reúne validações e normalizações de entrada reutilizáveis, sem dependência de domínio.
package validation

import (
	"errors"
	"net/mail"
	"strings"
	"unicode/utf8"
)

const (
	MaxTextLen     = 255
	MinPasswordLen = 8
	// O bcrypt rejeita entradas acima de 72 bytes
	MaxPasswordLen = 72
)

var (
	ErrInvalidName     = errors.New("nome inválido")
	ErrInvalidEmail    = errors.New("e-mail inválido")
	ErrInvalidPassword = errors.New("a senha deve ter entre 8 e 72 caracteres")
)

// Name remove espaços nas pontas e exige de 1 a MaxTextLen caracteres
func Name(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" || utf8.RuneCountInString(name) > MaxTextLen {
		return "", ErrInvalidName
	}
	return name, nil
}

// Email normaliza para minúsculas e valida o formato (apenas o endereço, sem nome de exibição)
func Email(email string) (string, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	addr, err := mail.ParseAddress(email)
	if err != nil || addr.Address != email || len(email) > MaxTextLen {
		return "", ErrInvalidEmail
	}
	return email, nil
}

// Password valida apenas o tamanho; não altera o valor
func Password(password string) error {
	if len(password) < MinPasswordLen || len(password) > MaxPasswordLen {
		return ErrInvalidPassword
	}
	return nil
}
