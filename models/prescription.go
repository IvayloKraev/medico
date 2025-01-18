package models

import (
	"github.com/google/uuid"
	"time"
)

type Prescription struct {
	ID                      uuid.UUID                `gorm:"primaryKey;unique;type:uuid;not null"`
	CreationDate            time.Time                `gorm:"not null"`
	DoctorID                uuid.UUID                `gorm:"type:uuid;not null"`
	Doctor                  Doctor                   `gorm:"foreignKey:DoctorID;references:ID"`
	CitizenID               uuid.UUID                `gorm:"type:uuid;not null"`
	PrescriptionMedicaments []PrescriptionMedicament `gorm:"foreignKey:PrescriptionID"`
	PrescriptionStateID     uuid.UUID                `gorm:"type:uuid;not null"`
	PrescriptionState       PrescriptionState        `gorm:"polymorphic:Owner;"`
}

type PrescriptionMedicament struct {
	PrescriptionID uuid.UUID  `gorm:"type:uuid;not null"`
	MedicamentID   uuid.UUID  `gorm:"type:uuid;not null"`
	Medicament     Medicament `gorm:"foreignKey:MedicamentID;references:ID"`
	Quantity       uint
}

type PrescriptionState struct {
	ID              uint   `gorm:"primaryKey"`
	OwnerID         uint   `gorm:"not null"`
	OwnerType       string `gorm:"not null"`
	State           string `gorm:"not null"`
	DeadlineDate    *time.Time
	FulfillmentDate *time.Time
	FulfilledByID   *uint
	ExpirationDate  *time.Time
}
