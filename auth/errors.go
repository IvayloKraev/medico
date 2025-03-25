package auth

import "errors"

const (
	WrongData = "wrong data"
)

const (
	EmailIncorrect = "the provided email is not correct"
)

var (
	ErrWrongData = errors.New(WrongData)
)

var (
	ErrEmailIncorrect = errors.New(EmailIncorrect)
)
