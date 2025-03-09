package dto

import (
	"errors"
	"github.com/google/uuid"
)

type RequestModeratorLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (m *RequestModeratorLogin) Validate() error {
	return errors.Join(
		validateEmail(m.Email),
		validateTotalNumberOfCharacters(m.Password))
}

type RequestModeratorCreateDoctor struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	LastName   string `json:"lastName"`
	UIN        string `json:"uin"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func (m *RequestModeratorCreateDoctor) Validate() error {
	return errors.Join(
		validateNameLength(m.FirstName, 3, 32),
		validateNameLength(m.SecondName, 3, 32),
		validateNameLength(m.LastName, 3, 32),
		validateEmail(m.Email),
		validateNumberOfLowerCase(m.Password),
		validateNumberOfUpperCase(m.Password),
		validateNumberOfDigits(m.Password),
		validateNumberOfSpecialCharacters(m.Password),
		validateTotalNumberOfCharacters(m.Password),
		validateNotIncludedWhiteSpaces(m.Password),
		validateUinLength(m.UIN))
}

type QueryModeratorDeleteDoctor struct {
	DoctorId uuid.UUID `json:"doctorId"`
}

type ResponseModeratorGetDoctors struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"firstName"`
	SecondName string    `json:"secondName"`
	LastName   string    `json:"lastName"`
	UIN        string    `json:"uin"`
	Email      string    `json:"email"`
}

type RequestModeratorCreateMedicament struct {
	OfficialName      string `json:"name"`
	ActiveIngredients []struct {
		Name string `json:"value"`
	} `json:"activeIngredients"`
	ATC string `json:"atc"`
}

func (m *RequestModeratorCreateMedicament) Validate() error {
	return errors.Join(
		validateNameLength(m.OfficialName, 3, 1000),
		validateAtcCode(m.ATC))
}

type QueryModeratorDeleteMedicament struct {
	MedicamentId uuid.UUID `json:"medicamentId"`
}
type ResponseModeratorGetMedicaments struct {
	ID                uuid.UUID `json:"id"`
	OfficialName      string    `json:"name"`
	ActiveIngredients []string  `json:"activeIngredients"`
	ATC               string    `json:"atc"`
}

type RequestModeratorCreatePharmacy struct {
	Name          string `json:"name"`
	OwnerName     string `json:"ownerName"`
	OwnerEmail    string `json:"ownerEmail"`
	OwnerPassword string `json:"ownerPassword"`
}

func (m *RequestModeratorCreatePharmacy) Validate() error {
	return errors.Join(
		validateNameLength(m.Name, 1, 300),
		validateNameLength(m.OwnerName, 3, 32),
		validateEmail(m.OwnerEmail),
		validateNumberOfLowerCase(m.OwnerPassword),
		validateNumberOfUpperCase(m.OwnerPassword),
		validateNumberOfDigits(m.OwnerPassword),
		validateNumberOfSpecialCharacters(m.OwnerPassword),
		validateTotalNumberOfCharacters(m.OwnerPassword),
		validateNotIncludedWhiteSpaces(m.OwnerPassword))
}

type QueryModeratorDeletePharmacy struct {
	PharmacyId uuid.UUID `json:"pharmacyId"`
}
type ResponseModeratorGetPharmacies struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	OwnerName string    `json:"ownerName"`
}

type RequestModeratorCreateCitizen struct {
	FirstName        string    `json:"firstName"`
	SecondName       string    `json:"secondName"`
	LastName         string    `json:"lastName"`
	UCN              string    `json:"ucn"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	PersonalDoctorId uuid.UUID `json:"personal_doctor_id"`
}

func (m *RequestModeratorCreateCitizen) Validate() error {
	return errors.Join(
		validateNameLength(m.FirstName, 3, 32),
		validateNameLength(m.SecondName, 3, 32),
		validateNameLength(m.LastName, 3, 32),
		validateEmail(m.Email),
		validateNumberOfLowerCase(m.Password),
		validateNumberOfUpperCase(m.Password),
		validateNumberOfDigits(m.Password),
		validateNumberOfSpecialCharacters(m.Password),
		validateTotalNumberOfCharacters(m.Password),
		validateNotIncludedWhiteSpaces(m.Password),
		validateUcn(m.UCN))
}

type QueryModeratorDeleteCitizen struct {
	CitizenId uuid.UUID `json:"citizenId"`
}
type ResponseModeratorGetCitizens struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"firstName"`
	SecondName string    `json:"secondName"`
	LastName   string    `json:"lastName"`
	UCN        string    `json:"ucn"`
}
