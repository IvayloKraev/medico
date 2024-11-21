package validators

import (
	"errors"
	"medico/errors"
	"regexp"
)

func ValidateEmail(email string) error {
	var emailRegexp = regexp.MustCompile(EmailPattern)

	if !emailRegexp.MatchString(email) {
		return errors.New(medicoErrors.IncorrectEmail)
	}
	return nil
}

func ValidateNumberOfLowerCase(password string) error {
	var lowerCaseRegexp = regexp.MustCompile(LowerCasePattern)

	if !lowerCaseRegexp.MatchString(password) {
		return errors.New(medicoErrors.NotEnoughLowerCase)
	}
	return nil
}

func ValidateNumberOfUpperCase(password string) error {
	var upperCaseRegexp = regexp.MustCompile(UpperCasePattern)

	if !upperCaseRegexp.MatchString(password) {
		return errors.New(medicoErrors.NotEnoughUpperCase)
	}
	return nil
}

func ValidateNumberOfDigits(password string) error {
	var digitsRegexp = regexp.MustCompile(DigitPattern)

	if !digitsRegexp.MatchString(password) {
		return errors.New(medicoErrors.NotEnoughDigits)
	}
	return nil
}

func ValidateNumberOfSpecialCharacters(password string) error {
	var specialCharsRegexp = regexp.MustCompile(SpecialCharPattern)

	if !specialCharsRegexp.MatchString(password) {
		return errors.New(medicoErrors.NotEnoughSpecialChars)
	}
	return nil
}

func ValidateTotalNumberOfCharacters(password string) error {
	var numberOfCharsRegexp = regexp.MustCompile(NumOfCharPattern)

	if !numberOfCharsRegexp.MatchString(password) {
		return errors.New(medicoErrors.NotEnoughNumberOfChars)
	}
	return nil
}

func ValidateNotIncludedWhiteSpaces(password string) error {
	var spaceRegexp = regexp.MustCompile(SpacePattern)

	if !spaceRegexp.MatchString(password) {
		return errors.New(medicoErrors.IncludedWhiteSpace)
	}
	return nil
}
