package dto

import (
	"errors"
	"medico/validators"
)

type CommonUserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CommonUserSignUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DoctorSignIn struct {
}

func (cusu CommonUserSignUp) Validate() error {
	return errors.Join(
		validators.ValidateEmail(cusu.Email),
		validators.ValidateTotalNumberOfCharacters(cusu.Password))
}

func (cusi CommonUserSignIn) Validate() error {
	return errors.Join(
		validators.ValidateEmail(cusi.Email),
		validators.ValidateNumberOfLowerCase(cusi.Password),
		validators.ValidateNumberOfUpperCase(cusi.Password),
		validators.ValidateNumberOfDigits(cusi.Password),
		validators.ValidateNumberOfSpecialCharacters(cusi.Password),
		validators.ValidateTotalNumberOfCharacters(cusi.Password),
		validators.ValidateNotIncludedWhiteSpaces(cusi.Password))
}
