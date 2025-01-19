package models

type Unit struct {
	ID   TModelID `gorm:"not null;type:uuid;primary_key"`
	Unit TShortText
}

type MedicamentApplication struct {
	ID          TModelID `gorm:"not null;type:uuid;primary_key"`
	Application TShortText
}

type AuthorizationHolder struct {
	ID      TModelID `gorm:"not null;type:uuid;primary_key"`
	Name    TShortText
	Country TShortText
}

type ActiveIngredient struct {
	ID                  TModelID `gorm:"not null;type:uuid;primary_key"`
	OfficialName        TShortText
	BulgarianName       TShortText
	MaximumDosage       float32
	MaximumDosageUnitID TModelID `gorm:"not null;type:uuid"`
	MaximumDosageUnit   Unit     `gorm:"foreignKey:MaximumDosageUnitID;references:ID"`
	Description         TLongtext
}

type ActiveIngredientInteraction struct {
	ActiveIngredient1ID TModelID         `gorm:"not null;type:uuid"`
	ActiveIngredient1   ActiveIngredient `gorm:"foreignKey:ActiveIngredient1ID;references:ID"`
	ActiveIngredient2ID TModelID         `gorm:"not null;type:uuid"`
	ActiveIngredient2   ActiveIngredient `gorm:"foreignKey:ActiveIngredient2ID;references:ID"`
	Description         TLongtext
}

type Medicament struct {
	ID                    TModelID `gorm:"not null;type:uuid;primary_key"`
	RegionalNumber        int
	Identification        TShortText
	OfficialName          TShortText
	BulgarianName         TShortText
	Description           TLongtext
	ActiveIngredients     []ActiveIngredient    `gorm:"many2many:active_ingredients_medicaments;"`
	ApplicationID         TModelID              `gorm:"not null;type:uuid"`
	Application           MedicamentApplication `gorm:"foreignKey:ApplicationID;references:ID"`
	Quantity              int
	UnitID                TModelID            `gorm:"not null;type:uuid"`
	Unit                  Unit                `gorm:"foreignKey:UnitID;references:ID"`
	AuthorizationHolderID TModelID            `gorm:"not null;type:uuid"`
	AuthorisationHolder   AuthorizationHolder `gorm:"foreignKey:AuthorizationHolderID;references:ID"`
	ATC                   TShortText
	RequiredPrescription  bool
}
