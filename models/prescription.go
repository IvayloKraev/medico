package models

import "github.com/google/uuid"

type PrescriptionState string

const (
	Active    PrescriptionState = "active"
	Invalid   PrescriptionState = "invalid"
	Fulfilled PrescriptionState = "fulfilled"
)

type Prescription struct {
	ID           uuid.UUID                `gorm:"primaryKey;unique;type:uuid;not null"`
	DoctorID     uuid.UUID                `gorm:"type:uuid;not null"`
	Doctor       Doctor                   `gorm:"foreignKey:DoctorID;references:ID"`
	CitizenID    uuid.UUID                `gorm:"type:uuid;not null"`
	Medicaments  []PrescriptionMedicament `gorm:"foreignKey:PrescriptionID"`
	State        PrescriptionState        `gorm:"type:enum('active','fulfilled','invalid'); not null"`
	Name         Text
	CreationDate DateTime `gorm:"not null"`
	StartDate    DateTime `gorm:"not null"`
	EndDate      DateTime `gorm:"not null"`
}

type PrescriptionMedicament struct {
	PrescriptionID uuid.UUID  `gorm:"primaryKey;type:uuid;not null"`
	MedicamentID   uuid.UUID  `gorm:"type:uuid;not null"`
	Medicament     Medicament `gorm:"foreignKey:MedicamentID;references:ID"`
	Quantity       uint
	Fulfilled      bool
}
