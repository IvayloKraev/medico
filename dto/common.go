package dto

import (
	"errors"
	"fmt"
	"medico/errors"
)

type Validate interface {
	Validate() error
}

type ToString interface {
	ToString() string
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
	result := validateEmail(email.ToString())

	if result != nil {
		return errors.New(fmt.Sprintf("%s - %s", medicoErrors.InvalidEmail, medicoErrors.IncorrectEmail))
	}

	return nil
}

func (password Password) Validate() error {
	lowerCaseResult := validateNumberOfLowerCase(password.ToString())
	upperCaseResult := validateNumberOfUpperCase(password.ToString())
	numberOfDigitsResult := validateNumberOfDigits(password.ToString())
	numberOfSpecialCharacters := validateNumberOfSpecialCharacters(password.ToString())
	numberOfCharacters := validateTotalNumberOfCharacters(password.ToString())
	notIncludedWhiteSpacesResult := validateNotIncludedWhiteSpaces(password.ToString())

	return errors.Join(
		lowerCaseResult,
		upperCaseResult,
		numberOfDigitsResult,
		numberOfSpecialCharacters,
		numberOfCharacters,
		notIncludedWhiteSpacesResult,
	)
}
