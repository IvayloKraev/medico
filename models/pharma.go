package models

import (
	"github.com/google/uuid"
	"time"
)

type PharmacyOwnerAuth struct {
	ID            uuid.UUID     `gorm:"primary_key;unique;type:uuid;not null"`
	Email         string        `gorm:"type:text;not null"`
	Password      string        `gorm:"type:text;not null"`
	PharmacyOwner PharmacyOwner `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE;"`
}

type PharmacyOwner struct {
	ID   uuid.UUID `gorm:"not null;primary_key;type:uuid;"`
	Name string
}

type PharmacyBrand struct {
	ID                  uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	Name                string
	Website             string
	OwnerID             uuid.UUID     `gorm:"not null;type:uuid"`
	Owner               PharmacyOwner `gorm:"foreignkey:OwnerID;references:ID"`
	HeadquartersAddress string
	PharmacyBranches    []PharmacyBranch `gorm:"foreignKey:PharmacyBrandID;"`
}

type PharmacyBranch struct {
	ID                uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	Address           string
	PharmacyBrandID   uuid.UUID     `gorm:"not null;type:uuid"`
	PharmacyBrand     PharmacyBrand `gorm:"foreignKey:PharmacyBrandID;references:ID"`
	Latitude          float32
	Longitude         float32
	Storage           []PharmacyBranchStorage `gorm:"foreignKey:PharmacyBranchID;"`
	Pharmacists       []Pharmacist            `gorm:"foreignKey:PharmacyBranchID;"`
	WorkdaysStartTime time.Time
	WorkdaysEndTime   time.Time
	WeekendsStartTime time.Time
	WeekendsEndTime   time.Time
}

type PharmacyBranchStorage struct {
	PharmacyBranchID uuid.UUID  `gorm:"not null;type:uuid"`
	MedicamentID     uuid.UUID  `gorm:"not null;type:uuid"`
	Medicament       Medicament `gorm:"foreignKey:MedicamentID;references:ID"`
	Quantity         WholeQuantity
}

type Pharmacist struct {
	ID               uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	FirstName        Text
	SecondName       Text
	Surname          Text
	PharmacyBranchID uuid.UUID      `gorm:"not null;type:uuid"`
	PharmacyBranch   PharmacyBranch `gorm:"foreignKey:PharmacyBranchID;"`
}
