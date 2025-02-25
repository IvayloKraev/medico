package models

import (
	"errors"
	"github.com/google/uuid"
)

type ModeratorType string

const (
	DoctorMod     ModeratorType = "doctorMod"
	CitizenMod                  = "citizenMod"
	PharmacyMod                 = "pharmacyMod"
	MedicamentMod               = "medicamentMod"
)

func ModeratorTypeFromText(text string) (ModeratorType, error) {
	modType := ModeratorType(text)

	if modType != DoctorMod && modType != CitizenMod && modType != MedicamentMod && modType != PharmacyMod {
		return "", errors.New("not logged in")
	}

	return modType, nil
}

type ModeratorAuth struct {
	ID        uuid.UUID `gorm:"primary_key;unique;type:uuid;not null"`
	Email     string    `gorm:"type:text;not null"`
	Password  string    `gorm:"type:text;not null"`
	Moderator Moderator `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE;"`
}

type Moderator struct {
	ID         uuid.UUID `gorm:"primary_key;type:uuid;not null"`
	FirstName  string
	SecondName string
	LastName   string
	Email      string
	Type       ModeratorType `gorm:"type:enum('doctorMod','citizenMod','pharmacyMod','medicamentMod');not null"`
}
