package dto

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type RequestDoctorLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (d *RequestDoctorLogin) Validate() error {
	return errors.Join(
		validateEmail(d.Email),
		validateTotalNumberOfCharacters(d.Password))
}

type RequestDoctorCreatePrescription struct {
	CitizenId   uuid.UUID `json:"citizen_id"`
	Name        string    `json:"name"`
	EndDate     time.Time `json:"end_date"`
	Medicaments []struct {
		OfficialName string `json:"official_name"`
		Quantity     uint   `json:"quantity"`
	} `json:"medicaments"`
}

func (d *RequestDoctorCreatePrescription) Validate() error {
	return errors.Join(
		validateNameLength(d.Name, 3, 32),
		validateTime(d.EndDate, time.Now(), TimeAfter))
}

type QueryDoctorGetCitizenInfo struct {
	CitizenUcn string `json:"citizen_ucn"`
}

type QueryDoctorGetCitizenPrescription struct {
	CitizenId uuid.UUID `json:"citizen_id"`
}

type ResponseDoctorCitizenInfo struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	LastName   string    `json:"last_name"`
	BirthDate  time.Time `json:"birth_date"`
	Email      string    `json:"email"`
}

type ResponseDoctorGetCitizenPrescriptionResponse struct {
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
