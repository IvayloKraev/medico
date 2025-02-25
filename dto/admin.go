package dto

import (
	"errors"
	"github.com/google/uuid"
	"medico/models"
)

type AdminLogin struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
}

func (a AdminLogin) Validate() error {
	return errors.Join(a.Email.Validate(), a.Password.Validate())
}

type AdminCreateModerator struct {
	FirstName  string               `json:"first_name"`
	SecondName string               `json:"second_name"`
	LastName   string               `json:"last_name"`
	Email      string               `json:"email"`
	Password   string               `json:"password"`
	Type       models.ModeratorType `json:"type"`
}

func (a AdminCreateModerator) Validate() error {
	if a.Type == "" {
		return errors.New("invalid role")
	}

	return nil
}

type AdminDeleteModerator struct {
	ModeratorId uuid.UUID `json:"moderatorId"`
}

type AdminGetModerator struct {
	ID         uuid.UUID            `json:"id"`
	FirstName  string               `json:"first_name"`
	SecondName string               `json:"second_name"`
	LastName   string               `json:"last_name"`
	Email      string               `json:"email"`
	Type       models.ModeratorType `json:"type"`
}
