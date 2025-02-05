package dto

import (
	"errors"
	"medico/errors"
	"regexp"
)

func validateEmail(email string) error {
	var emailRegexp = regexp.MustCompile(emailPattern)

	if !emailRegexp.MatchString(email) {
		return errors.New(medicoErrors.IncorrectEmail)
	}
	return nil
}

func validateNumberOfLowerCase(password string) error {
	var lowerCaseRegexp = regexp.MustCompile(lowerCasePattern)

	if !lowerCaseRegexp.MatchString(password) {
		return errors.New(medicoErrors.NotEnoughLowerCase)
	}
	return nil
}

func validateNumberOfUpperCase(password string) error {
	var upperCaseRegexp = regexp.MustCompile(upperCasePattern)

	if !upperCaseRegexp.MatchString(password) {
		return errors.New(medicoErrors.NotEnoughUpperCase)
	}
	return nil
}

func validateNumberOfDigits(password string) error {
	var digitsRegexp = regexp.MustCompile(digitPattern)

	if !digitsRegexp.MatchString(password) {
		return errors.New(medicoErrors.NotEnoughDigits)
	}
	return nil
}

func validateNumberOfSpecialCharacters(password string) error {
	var specialCharsRegexp = regexp.MustCompile(specialCharPattern)

	if !specialCharsRegexp.MatchString(password) {
		return errors.New(medicoErrors.NotEnoughSpecialChars)
	}
	return nil
}

func validateTotalNumberOfCharacters(password string) error {
	var numberOfCharsRegexp = regexp.MustCompile(numOfCharPattern)

	if !numberOfCharsRegexp.MatchString(password) {
		return errors.New(medicoErrors.NotEnoughNumberOfChars)
	}
	return nil
}

func validateNotIncludedWhiteSpaces(password string) error {
	var spaceRegexp = regexp.MustCompile(spacePattern)

	if !spaceRegexp.MatchString(password) {
		return errors.New(medicoErrors.IncludedWhiteSpace)
	}
	return nil
}
