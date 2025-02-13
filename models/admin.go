package models

import "github.com/google/uuid"

type AdminAuth struct {
	ID       uuid.UUID `gorm:"primary_key;unique;type:uuid;not null;"`
	Email    string    `gorm:"type:text;not null"`
	Password string    `gorm:"type:text;not null"`
}
