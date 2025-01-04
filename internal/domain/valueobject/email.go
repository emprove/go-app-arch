package valueobject

import (
	"errors"
	"go-app-arch/internal/validation"
)

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	if !validation.IsEmail(value) {
		return Email{}, errors.New("invalid email")
	}
	return Email{value: value}, nil
}

func (e Email) String() string {
	return e.value
}
