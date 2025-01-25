package models

import "time"

type PharmacyBrand struct {
	ID                  ModelID `gorm:"not null;type:uuid;primary_key"`
	Name                Text
	Website             Text
	Owner               Text
	HeadquartersAddress Text
	PharmacyBranches    []PharmacyBranch `gorm:"foreignKey:PharmacyBrandID;"`
}

type PharmacyBranch struct {
	ID                ModelID `gorm:"not null;type:uuid;primary_key"`
	Address           Text
	PharmacyBrandID   ModelID       `gorm:"not null;type:uuid"`
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
	PharmacyBranchID ModelID    `gorm:"not null;type:uuid"`
	MedicamentID     ModelID    `gorm:"not null;type:uuid"`
	Medicament       Medicament `gorm:"foreignKey:MedicamentID;references:ID"`
	Quantity         WholeQuantity
}

type Pharmacist struct {
	ID               ModelID `gorm:"not null;type:uuid;primary_key"`
	FirstName        Text
	SecondName       Text
	Surname          Text
	PharmacyBranchID ModelID        `gorm:"not null;type:uuid"`
	PharmacyBranch   PharmacyBranch `gorm:"foreignKey:PharmacyBranchID;"`
}
