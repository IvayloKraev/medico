package models

import "github.com/google/uuid"

type Hospital struct {
	ID      uuid.UUID `gorm:"primaryKey;unique;type:uuid;not null"`
	Name    Text
	Address Text
	Doctors []Doctor `gorm:"foreignKey:HospitalID"`
}

type DoctorAuth struct {
	ID       uuid.UUID `gorm:"primary_key;unique;type:uuid;not null"`
	Email    string    `gorm:"type:text;not null"`
	Password string    `gorm:"type:text;not null"`
	Doctor   Doctor    `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE;"`
}

type Doctor struct {
	ID         uuid.UUID `gorm:"primaryKey;unique;type:uuid;not null"`
	FirstName  string
	SecondName string
	LastName   string
	HospitalID uuid.UUID `gorm:"type:uuid;not null"`
	Hospital   Hospital  `gorm:"foreignKey:HospitalID"`
	UIN        string
	Email      string
}
