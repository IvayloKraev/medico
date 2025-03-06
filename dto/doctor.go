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
	CitizenId   uuid.UUID `json:"citizenId"`
	Name        string    `json:"name"`
	EndDate     time.Time `json:"end_date"`
	Medicaments []struct {
		Id       uuid.UUID `json:"id"`
		Quantity uint      `json:"quantity"`
	} `json:"medicaments"`
}

func (d *RequestDoctorCreatePrescription) Validate() error {
	return errors.Join(
		validateNameLength(d.Name, 3, 32),
		validateTime(d.EndDate, time.Now(), TimeAfter))
}

type QueryDoctorGetCitizenInfo struct {
	CitizenUcn string `json:"citizenUcn"`
}

type ResponseListOfCitizensViaCommonUCN struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UCN       string `json:"ucn"`
}

type QueryDoctorGetCitizenPrescription struct {
	CitizenId uuid.UUID `json:"citizenId"`
}

type ResponseDoctorCitizenInfo struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"firstName"`
	SecondName string    `json:"secondName"`
	LastName   string    `json:"lastName"`
	BirthDate  time.Time `json:"birthDate"`
	Email      string    `json:"email"`
}

type ResponseDoctorGetCitizenPrescription struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Medicaments []struct {
		OfficialName string `json:"officialName"`
		Quantity     uint   `json:"quantity"`
	} `json:"medicaments"`
	State       string    `json:"status"`
	CreatedDate time.Time `json:"createdDate"`
	StartDate   time.Time `json:"issuedDate"`
	EndDate     time.Time `json:"endDate"`
}

type QueryDoctorGetMedicamentByCommonName struct {
	CommonName string `json:"name"`
}

type ResponseDoctorGetMedicamentPrescription struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"officialName"`
}
