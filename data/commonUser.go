package data

import "github.com/google/uuid"

type CommonUserDB struct {
	ID           uuid.UUID `gorm:"primaryKey;unique;type:varchar(32);not null"`
	FirstName    string    `gorm:""`
	SecondName   string    `gorm:""`
	Email        string    `gorm:""`
	Password     string    `gorm:""`
	PasswordSalt string    `gorm:""`
}
