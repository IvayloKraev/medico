package dto

import (
	"github.com/google/uuid"
	"time"
)

type DoctorLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DoctorCreatePrescription struct {
	CitizenId   uuid.UUID `json:"citizen_id"`
	Name        string    `json:"name"`
	EndDate     time.Time `json:"end_date"`
	Medicaments []struct {
		OfficialName string `json:"official_name"`
		Quantity     uint   `json:"quantity"`
	} `json:"medicaments"`
}

type DoctorGetCitizenInfo struct {
	CitizenUcn string `json:"citizen_ucn"`
}

type DoctorGetCitizenPrescription struct {
	CitizenId uuid.UUID `json:"citizen_id"`
}

type DoctorCitizenInfo struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	LastName   string    `json:"last_name"`
	BirthDate  time.Time `json:"birth_date"`
	Email      string    `json:"email"`
}

type DoctorGetCitizenPrescriptionResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string
	Medicaments []struct {
		OfficialName string `json:"official_name"`
		Quantity     uint   `json:"quantity"`
	} `json:"medicaments"`
	State       string
	CreatedDate time.Time `json:"created_date"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
