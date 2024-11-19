package data

import "github.com/google/uuid"

type CommonUserDB struct {
	ID           uuid.UUID `gorm:"primaryKey;unique;type:uuid;not null"`
	FirstName    string    `gorm:""`
	SecondName   string    `gorm:""`
	LastName     string    `gorm:""`
	Email        string    `gorm:""`
	Password     string    `gorm:""`
	PasswordSalt string    `gorm:""`
}
