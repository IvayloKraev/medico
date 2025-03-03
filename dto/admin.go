package dto

import (
	"errors"
	"github.com/google/uuid"
	"medico/common"
)

type RequestAdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *RequestAdminLogin) Validate() error {
	return errors.Join(
		validateEmail(a.Email),
		validateTotalNumberOfCharacters(a.Password))
}

type RequestAdminCreateModerator struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Type       string `json:"type"`
}

func (a *RequestAdminCreateModerator) Validate() error {
	return errors.Join(
		validateNameLength(a.FirstName, 3, 32),
		validateNameLength(a.SecondName, 3, 32),
		validateNameLength(a.LastName, 3, 32),
		validateEmail(a.Email),
		validateNumberOfLowerCase(a.Password),
		validateNumberOfUpperCase(a.Password),
		validateNumberOfDigits(a.Password),
		validateNumberOfSpecialCharacters(a.Password),
		validateTotalNumberOfCharacters(a.Password),
		validateNotIncludedWhiteSpaces(a.Password),
		validateModeratorType(a.Type))
}

type QueryAdminDeleteModerator struct {
	ModeratorId uuid.UUID `query:"moderatorId"`
}

type ResponseAdminGetModerator struct {
	ID         uuid.UUID            `json:"id"`
	FirstName  string               `json:"firstName"`
	SecondName string               `json:"secondName"`
	LastName   string               `json:"lastName"`
	Email      string               `json:"email"`
	Type       common.ModeratorType `json:"type"`
}
