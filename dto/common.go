package dto

import (
	"errors"
	"fmt"
	"medico/errors"
	"medico/validators"
)

type IValidate interface {
	Validate() error
}

type Email string
type Password string

func (email *Email) Validate() error {
	result := validators.ValidateEmail(string(*email))

	if result != nil {
		return errors.New(fmt.Sprintf("%s - %s", medicoErrors.InvalidEmail, medicoErrors.IncorrectEmail))
	}

	return nil
}

func (password *Password) Validate() error {
	lowerCaseResult := validators.ValidateNumberOfLowerCase(string(*password))

	if lowerCaseResult != nil {
	}

	return errors.ErrUnsupported
}
