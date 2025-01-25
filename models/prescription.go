package models

type PrescriptionState string

const (
	Active    PrescriptionState = "active"
	Invalid   PrescriptionState = "invalid"
	Fulfilled PrescriptionState = "fulfilled"
)

type Prescription struct {
	ID           ModelID                  `gorm:"primaryKey;unique;type:uuid;not null"`
	DoctorID     ModelID                  `gorm:"type:uuid;not null"`
	Doctor       Doctor                   `gorm:"foreignKey:DoctorID;references:ID"`
	CitizenID    ModelID                  `gorm:"type:uuid;not null"`
	Medicaments  []PrescriptionMedicament `gorm:"foreignKey:PrescriptionID"`
	State        PrescriptionState        `gorm:"type:enum('active','fulfilled','invalid'); not null"`
	Name         Text
	CreationDate DateTime `gorm:"not null"`
	StartDate    DateTime `gorm:"not null"`
	EndDate      DateTime `gorm:"not null"`
}

type PrescriptionMedicament struct {
	PrescriptionID ModelID    `gorm:"primaryKey;type:uuid;not null"`
	MedicamentID   ModelID    `gorm:"type:uuid;not null"`
	Medicament     Medicament `gorm:"foreignKey:MedicamentID;references:ID"`
	Quantity       uint
	Fulfilled      bool
}
