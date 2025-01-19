package models

type PharmacyBrand struct {
	ID                  TModelID `gorm:"not null;type:uuid;primary_key"`
	Name                TShortText
	Website             TShortText
	Owner               TShortText
	HeadquartersAddress TShortText
	PharmacyBranches    []PharmacyBranch `gorm:"foreignKey:PharmacyBrandID;"`
}

type PharmacyBranch struct {
	ID              TModelID `gorm:"not null;type:uuid;primary_key"`
	Address         TShortText
	PharmacyBrandID TModelID      `gorm:"not null;type:uuid"`
	PharmacyBrand   PharmacyBrand `gorm:"foreignKey:PharmacyBrandID;references:ID"`
	Latitude        float32
	Longitude       float32
	Storage         []PharmacyBranchStorage `gorm:"foreignKey:PharmacyBranchID;"`
	Pharmacists     []Pharmacist            `gorm:"foreignKey:PharmacyBranchID;"`
}
type PharmacyBranchStorage struct {
	PharmacyBranchID TModelID   `gorm:"not null;type:uuid"`
	MedicamentID     TModelID   `gorm:"not null;type:uuid"`
	Medicament       Medicament `gorm:"foreignKey:MedicamentID;references:ID"`
	Quantity         int
}

type Pharmacist struct {
	ID               TModelID `gorm:"not null;type:uuid;primary_key"`
	FirstName        TShortText
	SecondName       TShortText
	Surname          TShortText
	PharmacyBranchID TModelID       `gorm:"not null;type:uuid"`
	PharmacyBranch   PharmacyBranch `gorm:"foreignKey:PharmacyBranchID;"`
}
