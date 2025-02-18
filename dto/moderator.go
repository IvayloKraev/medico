package dto

import (
	"errors"
	"github.com/google/uuid"
)

type ModeratorLogin struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
}

func (m *ModeratorLogin) Validate() error {
	return errors.Join(m.Email.Validate(), m.Password.Validate())
}

type ModeratorCreateDoctor struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	LastName   string `json:"last_name"`
	UIN        string `json:"uin"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type ModeratorDeleteDoctor struct {
	DoctorId uuid.UUID `json:"doctor_id"`
}

type ModeratorGetDoctors struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	LastName   string    `json:"last_name"`
	UIN        string    `json:"uin"`
	Email      string    `json:"email"`
}

type ModeratorCreateMedicament struct {
	OfficialName      string   `json:"official_name"`
	ActiveIngredients []string `json:"active_ingredients"`
	ATC               string   `json:"atc"`
	//RequiredPrescription bool     `json:"required_prescriptions"`
}
type ModeratorDeleteMedicament struct {
	MedicamentId uuid.UUID `json:"medicament_id"`
}
type ModeratorGetMedicaments struct {
	ID                uuid.UUID `json:"id"`
	OfficialName      string    `json:"official_name"`
	ActiveIngredients []string  `json:"active_ingredients"`
	ATC               string    `json:"atc"`
}

//type ModeratorCreatePharmacyOwner struct {
//}
//type ModeratorDeletePharmacyOwner struct{}

type ModeratorCreatePharmacy struct {
	Name          string `json:"name"`
	OwnerName     string `json:"pharmacy_owner"`
	OwnerEmail    string `json:"owner_email"`
	OwnerPassword string `json:"owner_password"`
}
type ModeratorDeletePharmacy struct {
	PharmacyId uuid.UUID `json:"pharmacy_id"`
}
type ModeratorGetPharmacies struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	OwnerName string    `json:"pharmacy_owner"`
}

type ModeratorCreateCitizen struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	LastName   string `json:"last_name"`
	UCN        string `json:"ucn"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
type ModeratorDeleteCitizen struct {
	CitizenId uuid.UUID `json:"citizen_id"`
}
type ModeratorGetCitizens struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	LastName   string    `json:"last_name"`
	UCN        string    `json:"ucn"`
}
