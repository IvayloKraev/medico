package models

import (
	"github.com/google/uuid"
	"medico/common"
)

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
	Type       common.ModeratorType `gorm:"type:enum('doctor','citizen','pharmacy','medicament');not null"`
}
