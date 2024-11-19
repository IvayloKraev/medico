package dto

import (
	"errors"
	"medico/utils"
	"regexp"
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
	emailRegex, numOfCharRegexp := regexp.MustCompile(utils.EmailPattern), regexp.MustCompile(utils.NumOfCharPattern)

	if emailRegex.MatchString(cusu.Email) == false {
		return errors.New("invalid email")
	}

	if numOfCharRegexp.MatchString(cusu.Password) == false {
		return errors.New("invalid password")
	}

	return nil
}

func (cusi CommonUserSignIn) Validate() error {

	var (
		emailRegexp       = regexp.MustCompile(utils.EmailPattern)
		lowerCaseRegexp   = regexp.MustCompile(utils.LowerCasePattern)
		upperCaseRegexp   = regexp.MustCompile(utils.UpperCasePattern)
		digitRegexp       = regexp.MustCompile(utils.DigitPattern)
		specialCharRegexp = regexp.MustCompile(utils.SpecialCharPattern)
		numOfCharRegexp   = regexp.MustCompile(utils.NumOfCharPattern)
		spaceRegexp       = regexp.MustCompile(utils.SpacePattern)
	)

	var (
		emailError       error = nil
		lowerCaseError   error = nil
		upperCaseError   error = nil
		digitError       error = nil
		specialCharError error = nil
		numOfError       error = nil
		spaceError       error = nil
	)

	if emailRegexp.MatchString(cusi.Email) == false {
		emailError = errors.New("ERROR: INVALID EMAIL - The provided email is not correct")
	}

	if !lowerCaseRegexp.MatchString(cusi.Password) {
		lowerCaseError = errors.New("ERROR: INVALID PASSWORD - password must contain at least 2 lower case letters")
	}
	if !upperCaseRegexp.MatchString(cusi.Password) {
		upperCaseError = errors.New("ERROR: INVALID PASSWORD - password must contain at least 2 upper case letters")
	}
	if !digitRegexp.MatchString(cusi.Password) {
		digitError = errors.New("ERROR: INVALID PASSWORD - password must contain at least 2 digits")
	}
	if !specialCharRegexp.MatchString(cusi.Password) {
		specialCharError = errors.New("ERROR: INVALID PASSWORD - password must contain at least 2 special characters")
	}
	if !numOfCharRegexp.MatchString(cusi.Password) {
		numOfError = errors.New("ERROR: INVALID PASSWORD - password must be at least 12 characters")
	}
	if spaceRegexp.MatchString(cusi.Password) {
		spaceError = errors.New("ERROR: INVALID PASSWORD - password must not include spaces")
	}

	return errors.Join(emailError, lowerCaseError, upperCaseError, digitError, specialCharError, numOfError, spaceError)
}
