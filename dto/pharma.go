package dto

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type RequestPharmacyOwnerAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *RequestPharmacyOwnerAuth) Validate() error {
	return errors.Join(
		validateEmail(p.Email),
		validateTotalNumberOfCharacters(p.Password))
}

type ResponsePharmacyOwnerBranches struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type RequestPharmacyOwnerNewBranch struct {
	Name      string  `json:"name"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func (p *RequestPharmacyOwnerNewBranch) Validate() error {
	return errors.Join(
		validateNameLength(p.Name, 3, 64),
		validateCoordinates(p.Latitude, p.Longitude))
}

type ResponsePharmacyOwnerPharmacist struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type RequestPharmacyOwnerNewPharmacist struct {
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	WorkingBranch uuid.UUID `json:"working_branch"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
}

func (p *RequestPharmacyOwnerNewPharmacist) Validate() error {
	return errors.Join(
		validateNameLength(p.FirstName, 3, 32),
		validateNameLength(p.LastName, 3, 32),
		validateNameLength(p.FirstName, 3, 64),
		validateNumberOfLowerCase(p.Password),
		validateNumberOfUpperCase(p.Password),
		validateNumberOfDigits(p.Password),
		validateNumberOfSpecialCharacters(p.Password),
		validateTotalNumberOfCharacters(p.Password),
		validateNotIncludedWhiteSpaces(p.Password))
}

type RequestPharmacistAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *RequestPharmacistAuth) Validate() error {
	return errors.Join(
		validateEmail(p.Email),
		validateTotalNumberOfCharacters(p.Password))
}

type PharmacistCitizenPrescriptionGet struct {
	CitizenUCN string `json:"citizen_ucn"`
}

type RequestPharmacistCitizenFulfillWholePrescription struct {
	CitizenId      uuid.UUID `json:"citizen_id"`
	BranchId       uuid.UUID `json:"branch_id"`
	PrescriptionId uuid.UUID `json:"prescription_id"`
}

type RequestPharmacistCitizenFulfillMedicamentFromPrescription struct {
	CitizenId      uuid.UUID `json:"citizen_ucn"`
	PrescriptionId uuid.UUID `json:"prescription_id"`
	MedicamentId   uuid.UUID `json:"medicament_ucn"`
}

type RequestPharmacistBranchAddMedicament struct {
	BranchId    uuid.UUID `json:"branch_id"`
	Medicaments []struct {
		BranchId       uuid.UUID `json:"branch_id"`
		MedicamentName string    `json:"medicament_name"`
		Quantity       uint      `json:"quantity"`
	} `json:"medicaments"`
}

type ResponsePharmacistCitizenPrescription struct {
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creation_date"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	Medicaments  []struct {
		MedicamentName string `json:"medicament_name"`
		Quantity       uint   `json:"quantity"`
		Fulfilled      bool   `json:"fulfilled"`
	} `json:"medicaments"`
}
