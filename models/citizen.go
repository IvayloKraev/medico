package models

import "github.com/google/uuid"

type Citizen struct {
	ID               uuid.UUID `gorm:"primaryKey;unique;type:uuid;not null"`
	FirstName        string
	SecondName       string
	Surname          string
	Age              int
	Height           float32
	Weight           float32
	Email            string
	Address          string
	City             string
	PersonalDoctorID uuid.UUID      `gorm:"type:uuid;not null"`
	PersonalDoctor   Doctor         `gorm:"foreignKey:PersonalDoctorID;references:ID"`
	Prescriptions    []Prescription `gorm:"foreignKey:CitizenID"`
}
