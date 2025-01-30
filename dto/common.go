package dto

import (
	"errors"
	"fmt"
	"medico/dto/validators"
	"medico/errors"
)

type Validate interface {
	Validate() error
}

type ToString interface {
	String() string
}

type Email string
type Password string

func (password Password) ToString() string {
	return string(password)
}

func (email Email) ToString() string {
	return string(email)
}

func (email Email) Validate() error {
	result := validators.ValidateEmail(email.ToString())

	if result != nil {
		return errors.New(fmt.Sprintf("%s - %s", medicoErrors.InvalidEmail, medicoErrors.IncorrectEmail))
	}

	return nil
}

func (password Password) Validate() error {
	lowerCaseResult := validators.ValidateNumberOfLowerCase(password.ToString())

	if lowerCaseResult != nil {
	}

	return errors.ErrUnsupported
}
