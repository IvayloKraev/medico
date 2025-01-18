package models

import (
	"github.com/google/uuid"
)

type Unit struct {
	ID   uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	Unit string
}

type MedicamentApplication struct {
	ID          uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	Application string
}

type AuthorizationHolder struct {
	ID      uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	Name    string
	Country string
}

type ActiveIngredient struct {
	ID                  uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	OfficialName        string
	BulgarianName       string
	MaximumDosage       float32
	MaximumDosageUnitID uuid.UUID `gorm:"not null;type:uuid"`
	MaximumDosageUnit   Unit      `gorm:"foreignKey:MaximumDosageUnitID;references:ID"`
	Description         string
}

type ActiveIngredientInteraction struct {
	ActiveIngredient1ID uuid.UUID        `gorm:"not null;type:uuid"`
	ActiveIngredient1   ActiveIngredient `gorm:"foreignKey:ActiveIngredient1ID;references:ID"`
	ActiveIngredient2ID uuid.UUID        `gorm:"not null;type:uuid"`
	ActiveIngredient2   ActiveIngredient `gorm:"foreignKey:ActiveIngredient2ID;references:ID"`
	Description         string
}

type Medicament struct {
	ID                    uuid.UUID `gorm:"not null;type:uuid;primary_key"`
	RegionalNumber        int
	Identification        string
	OfficialName          string
	BulgarianName         string
	Description           string
	ActiveIngredients     []ActiveIngredient    `gorm:"many2many:active_ingredients_medicaments;"`
	ApplicationID         uuid.UUID             `gorm:"not null;type:uuid"`
	Application           MedicamentApplication `gorm:"foreignKey:ApplicationID;references:ID"`
	Quantity              int
	UnitID                uuid.UUID           `gorm:"not null;type:uuid"`
	Unit                  Unit                `gorm:"foreignKey:UnitID;references:ID"`
	AuthorizationHolderID uuid.UUID           `gorm:"not null;type:uuid"`
	AuthorisationHolder   AuthorizationHolder `gorm:"foreignKey:AuthorizationHolderID;references:ID"`
	ATC                   string
	RequiredPrescription  bool
}
