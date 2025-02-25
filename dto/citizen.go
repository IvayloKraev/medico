package dto

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type CitizenLogin struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
}

func (l CitizenLogin) Validate() error {
	return errors.Join(l.Email.Validate(), l.Password.Validate())
}

type CitizenMedicalInfo struct {
	FirstName      string    `json:"first_name"`
	SecondName     string    `json:"second_name"`
	LastName       string    `json:"last_name"`
	BirthDate      time.Time `json:"birth_date"`
	Sex            string    `json:"sex"`
	UCN            string    `json:"ucn"`
	PersonalDoctor struct {
		FirstName  string `json:"first_name"`
		SecondName string `json:"second_name"`
		LastName   string `json:"last_name"`
		UIN        string `json:"uin"`
		Email      string `json:"email"`
	} `json:"personal_doctor"`
}

type CitizenAvailablePharmacyGet struct {
	PrescriptionId uuid.UUID `json:"prescription_id"`
}

type CitizenAvailablePharmacy struct {
	Name      string  `json:"name"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type CitizenPrescription struct {
	Doctor struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		UIN       string `json:"uin"`
	} `json:"doctor"`
	Medicaments []struct {
		Name string `json:"name"`
		Unit uint   `json:"unit"`
	} `json:"medicaments"`
}
