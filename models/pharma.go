package models

import "github.com/google/uuid"

type PharmacyBrand struct {
	ID                  uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	Name                string
	Website             string
	Owner               string
	HeadquartersAddress string
	PharmacyBranches    []PharmacyBranch `gorm:"foreignKey:PharmacyBrandID;"`
}

type PharmacyBranch struct {
	ID              uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	Address         string
	PharmacyBrandID uuid.UUID     `gorm:"not null;type:uuid"`
	PharmacyBrand   PharmacyBrand `gorm:"foreignKey:PharmacyBrandID;references:ID"`
	Latitude        float32
	Longitude       float32
	Storage         []PharmacyBranchStorage `gorm:"foreignKey:PharmacyBranchID;"`
	Pharmacists     []Pharmacist            `gorm:"foreignKey:PharmacyBranchID;"`
}
type PharmacyBranchStorage struct {
	PharmacyBranchID uuid.UUID  `gorm:"not null;type:uuid"`
	MedicamentID     uuid.UUID  `gorm:"not null;type:uuid"`
	Medicament       Medicament `gorm:"foreignKey:MedicamentID;references:ID"`
	Quantity         int
}

type Pharmacist struct {
	ID               uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	FirstName        string
	SecondName       string
	Surname          string
	PharmacyBranchID uuid.Time      `gorm:"not null;type:uuid"`
	PharmacyBranch   PharmacyBranch `gorm:"foreignKey:PharmacyBranchID;"`
}
