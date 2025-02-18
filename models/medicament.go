package models

import "github.com/google/uuid"

type Unit string

const (
	Milligrams Unit = "mg"
	Grams      Unit = "g"
)

type MedicamentApplication string

const (
	HardTablets MedicamentApplication = "hard_tablets"
	SoftTables  MedicamentApplication = "soft_tables"
)

type AuthorizationHolder struct {
	ID      uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	Name    Text
	Country Text
}

type ActiveIngredient struct {
	ID            uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	OfficialName  Text
	BulgarianName Text
	Description   Text
}

type ActiveIngredientInteraction struct {
	ActiveIngredient1ID uuid.UUID        `gorm:"not null;type:uuid"`
	ActiveIngredient1   ActiveIngredient `gorm:"foreignKey:ActiveIngredient1ID;references:ID"`
	ActiveIngredient2ID uuid.UUID        `gorm:"not null;type:uuid"`
	ActiveIngredient2   ActiveIngredient `gorm:"foreignKey:ActiveIngredient2ID;references:ID"`
	Description         Text
}

type ActiveIngredientsMedicament struct {
	ActiveIngredientID uuid.UUID        `gorm:"primary_key;not null;type:uuid"`
	ActiveIngredient   ActiveIngredient `gorm:"foreignKey:ActiveIngredientID;references:ID"`
	Quantity           WholeQuantity
	Unit               Unit `gorm:"type:enum('mg','g');"`
}

type Medicament struct {
	ID             uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	RegionalNumber int
	Identification string
	OfficialName   string
	BulgarianName  string
	Description    string
	//ActiveIngredients     []ActiveIngredientsMedicament `gorm:"many2many:active_ingredients_medicament;"`
	ActiveIngredients     string
	Application           MedicamentApplication `gorm:"type:enum('hard_tablets','soft_tables');"`
	ApplicationQuantity   int
	ApplicationUnit       Unit                `gorm:"foreignKey:UnitID;references:ID"`
	AuthorizationHolderID uuid.UUID           `gorm:"not null;type:uuid"`
	AuthorisationHolder   AuthorizationHolder `gorm:"foreignKey:AuthorizationHolderID;references:ID"`
	ATC                   string
	RequiredPrescription  bool
}
