package dto

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type RequestCitizenLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *RequestCitizenLogin) Validate() error {
	return errors.Join(
		validateEmail(c.Email),
		validateTotalNumberOfCharacters(c.Password))
}

type ResponseCitizenMedicalInfo struct {
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	LastName   string    `json:"last_name"`
	BirthDate  time.Time `json:"birth_date"`
	Sex        string    `json:"sex"`
	UCN        string    `json:"ucn"`
}

type ResponseCitizenPersonalDoctor struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	LastName   string `json:"last_name"`
	UIN        string `json:"uin"`
	Email      string `json:"email"`
}

type QueryCitizenAvailablePharmacyGet struct {
	PrescriptionId uuid.UUID `json:"prescription_id"`
}

type ResponseCitizenAvailablePharmacy struct {
	Name      string  `json:"name"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type ResponseCitizenPrescription struct {
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
