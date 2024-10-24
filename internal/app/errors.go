package app

import (
	"go-app-arch/internal/validation"
	"strings"
)

type ValidationError struct {
	Validator validation.Validator
}

func (e *ValidationError) Error() string {
	return strings.Join(e.Validator.Errors, ";")
}
