package models

import (
	"time"
)

type Prescription struct {
	ID                      TModelID                 `gorm:"primaryKey;unique;type:uuid;not null"`
	CreationDate            time.Time                `gorm:"not null"`
	DoctorID                TModelID                 `gorm:"type:uuid;not null"`
	Doctor                  Doctor                   `gorm:"foreignKey:DoctorID;references:ID"`
	CitizenID               TModelID                 `gorm:"type:uuid;not null"`
	PrescriptionMedicaments []PrescriptionMedicament `gorm:"foreignKey:PrescriptionID"`
	PrescriptionStateID     TModelID                 `gorm:"type:uuid;not null"`
	PrescriptionState       PrescriptionState        `gorm:"polymorphic:Owner;"`
}

type PrescriptionMedicament struct {
	PrescriptionID TModelID   `gorm:"type:uuid;not null"`
	MedicamentID   TModelID   `gorm:"type:uuid;not null"`
	Medicament     Medicament `gorm:"foreignKey:MedicamentID;references:ID"`
	Quantity       uint
}

type PrescriptionState struct {
	ID              uint       `gorm:"primaryKey"`
	OwnerID         uint       `gorm:"not null"`
	OwnerType       TShortText `gorm:"not null"`
	State           TShortText `gorm:"not null"`
	DeadlineDate    *time.Time
	FulfillmentDate *time.Time
	FulfilledByID   *uint
	ExpirationDate  *time.Time
}
