package dto

import (
	"medico/common"
	"regexp"
	"time"
)

type Validate interface {
	Validate() error
}

func validateEmail(email string) error {
	if !regexp.MustCompile(emailPattern).MatchString(email) {
		return ErrEmailIncorrect
	}
	return nil
}

func validateNumberOfLowerCase(password string) error {
	if !regexp.MustCompile(lowerCasePattern).MatchString(password) {
		return ErrPasswordNotEnoughLowerCase
	}
	return nil
}

func validateNumberOfUpperCase(password string) error {
	if !regexp.MustCompile(upperCasePattern).MatchString(password) {
		return ErrPasswordNotEnoughUpperCase
	}
	return nil
}

func validateNumberOfDigits(password string) error {
	if !regexp.MustCompile(digitPattern).MatchString(password) {
		return ErrPasswordNotEnoughDigits
	}
	return nil
}

func validateNumberOfSpecialCharacters(password string) error {
	if !regexp.MustCompile(specialCharPattern).MatchString(password) {
		return ErrPasswordNotEnoughSpecialChars
	}
	return nil
}

func validateTotalNumberOfCharacters(password string) error {
	if len(password) < 12 || len(password) > 72 {
		return ErrPasswordInvalidNumberOfChars
	}
	return nil
}

func validateNotIncludedWhiteSpaces(password string) error {
	if regexp.MustCompile(spacePattern).MatchString(password) {
		return ErrPasswordIncludedWhiteSpace
	}
	return nil
}

func validateNameLength(name string, min, max int) error {
	if len(name) < min || len(name) > max {
		return ErrNameInvalidNumberOfChars
	}

	return nil
}

func validateModeratorType(moderatorType string) error {
	if moderatorType != string(common.DoctorMod) &&
		moderatorType != string(common.PharmacyMod) &&
		moderatorType != string(common.MedicamentMod) &&
		moderatorType != string(common.CitizenMod) {
		return ErrModeratorTypeInvalid
	}
	return nil
}

const (
	TimeBefore = "timeBefore"
	TimeAfter  = "timeAfter"
)

func validateTime(time time.Time, pivot time.Time, comparison string) error {
	if comparison == TimeAfter && !time.After(pivot) {
		return ErrTimeInvalid
	} else if comparison == TimeBefore && !time.Before(pivot) {
		return ErrTimeInvalid
	}

	return nil
}

func validateUinLength(uin string) error {
	if len(uin) != 10 {
		return ErrUinInvalidLength
	}

	return nil
}

func validateAtcCode(atc string) error {
	if !regexp.MustCompile(atcPattern).MatchString(atc) {
		return ErrAtcInvalidCode
	}
	return nil
}

func validateUcn(ucn string) error {
	if len(ucn) != 10 {
		return ErrUcnInvalid
	}

	return nil
}

func validateCoordinates(latitude, longitude float32) error {
	if latitude < -90 && latitude > 90 && longitude < 0 && longitude > 180 {
		return ErrCoordinatesInvalid
	}
	return nil
}
