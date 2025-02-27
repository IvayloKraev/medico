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

type ModeratorCreateDoctor struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	LastName   string `json:"last_name"`
	UIN        string `json:"uin"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func (m *ModeratorCreateDoctor) Validate() error {
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

type ModeratorDeleteDoctor struct { // TODO: Make to param
	DoctorId uuid.UUID `json:"doctor_id"`
}

type ResponseModeratorGetDoctors struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	LastName   string    `json:"last_name"`
	UIN        string    `json:"uin"`
	Email      string    `json:"email"`
}

type RequestModeratorCreateMedicament struct {
	OfficialName      string   `json:"official_name"`
	ActiveIngredients []string `json:"active_ingredients"`
	ATC               string   `json:"atc"`
}

func (m *RequestModeratorCreateMedicament) Validate() error {
	return errors.Join(
		validateNameLength(m.OfficialName, 3, 1000),
		validateAtcCode(m.ATC))
}

type ModeratorDeleteMedicament struct { // TODO: Make to param
	MedicamentId uuid.UUID `json:"medicament_id"`
}
type ResponseModeratorGetMedicaments struct {
	ID                uuid.UUID `json:"id"`
	OfficialName      string    `json:"official_name"`
	ActiveIngredients []string  `json:"active_ingredients"`
	ATC               string    `json:"atc"`
}

type RequestModeratorCreatePharmacy struct {
	Name          string `json:"name"`
	OwnerName     string `json:"owner_name"`
	OwnerEmail    string `json:"owner_email"`
	OwnerPassword string `json:"owner_password"`
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

type ModeratorDeletePharmacy struct { // TODO: Make to param
	PharmacyId uuid.UUID `json:"pharmacy_id"`
}
type ResponseModeratorGetPharmacies struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	OwnerName string    `json:"pharmacy_owner"`
}

type RequestModeratorCreateCitizen struct {
	FirstName        string    `json:"first_name"`
	SecondName       string    `json:"second_name"`
	LastName         string    `json:"last_name"`
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

type ModeratorDeleteCitizen struct { // TODO: Make to param
	CitizenId uuid.UUID `json:"citizen_id"`
}
type ResponseModeratorGetCitizens struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	LastName   string    `json:"last_name"`
	UCN        string    `json:"ucn"`
}
