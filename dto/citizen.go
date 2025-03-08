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
	FirstName  string    `json:"firstName"`
	SecondName string    `json:"secondName"`
	LastName   string    `json:"lastName"`
	BirthDate  time.Time `json:"birthDate"`
	Sex        string    `json:"sex"`
	UCN        string    `json:"ucn"`
	Email      string    `json:"email"`
}

type ResponseCitizenPersonalDoctor struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	LastName   string `json:"last_name"`
	UIN        string `json:"uin"`
	Email      string `json:"email"`
}

type QueryCitizenAvailablePharmacyGet struct {
	PrescriptionId uuid.UUID `query:"prescriptionId"`
}

type ResponseCitizenAvailablePharmacy struct {
	Name      string  `json:"name"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type ResponseCitizenPrescription struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	State     string    `json:"status"`
	StartDate time.Time `json:"issuedDate"`
	Doctor    struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		UIN       string `json:"uin"`
	} `json:"doctor"`
	Medicaments []struct {
		Name     string `json:"officialName"`
		Quantity uint   `json:"quantity"`
	} `json:"medicaments"`
}
