package dto

import "errors"

const (
	WrongData = "wrong data"
)

const (
	EmailIncorrect = "the provided email is not correct"
)

const (
	PasswordNotEnoughLowerCase    = "password must contain at least 2 lower case letters"
	PasswordNotEnoughUpperCase    = "password must contain at least 2 upper case letters"
	PasswordNotEnoughDigits       = "password must contain at least 2 digits"
	PasswordNotEnoughSpecialChars = "password must contain at least 2 special characters"
	PasswordInvalidNumberOfChars  = "password must contain between 12 and 72 characters"
	PasswordIncludedWhiteSpace    = "password must not include spaces"
)

const (
	NameInvalidNumberOfChars = "name must contain between 3 and 32 characters"
)

const (
	ModeratorTypeInvalid = "provided moderator type is not valid"
)

const (
	TimeInvalid = "time is invalid"
)

const (
	UinInvalidLength = "uin of a doctor should be 10 characters long"
)

const (
	AtcInvalidCode = "atc code is invalid"
)

const (
	UcnInvalid = "ucn of citizen is invalid"
)

const (
	CoordinatesInvalid = "coordinates are invalid"
)

var (
	ErrWrongData = errors.New(WrongData)
)

var (
	ErrEmailIncorrect = errors.New(EmailIncorrect)
)

var (
	ErrPasswordNotEnoughLowerCase    = errors.New(PasswordNotEnoughLowerCase)
	ErrPasswordNotEnoughUpperCase    = errors.New(PasswordNotEnoughUpperCase)
	ErrPasswordNotEnoughDigits       = errors.New(PasswordNotEnoughDigits)
	ErrPasswordNotEnoughSpecialChars = errors.New(PasswordNotEnoughSpecialChars)
	ErrPasswordInvalidNumberOfChars  = errors.New(PasswordInvalidNumberOfChars)
	ErrPasswordIncludedWhiteSpace    = errors.New(PasswordIncludedWhiteSpace)
)

var (
	ErrNameInvalidNumberOfChars = errors.New(NameInvalidNumberOfChars)
)

var (
	ErrModeratorTypeInvalid = errors.New(ModeratorTypeInvalid)
)

var (
	ErrTimeInvalid = errors.New(TimeInvalid)
)

var (
	ErrUinInvalidLength = errors.New(UinInvalidLength)
)

var (
	ErrAtcInvalidCode = errors.New(AtcInvalidCode)
)

var (
	ErrUcnInvalid = errors.New(UcnInvalid)
)

var (
	ErrCoordinatesInvalid = errors.New(CoordinatesInvalid)
)
