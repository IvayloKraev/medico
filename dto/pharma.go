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

type QueryGetBranchesByCommonName struct {
	Name string `query:"name"`
}

type ResponseGetBranchesByCommonName struct {
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
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	WorkingBranch uuid.UUID `json:"pharmacy"`
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

type QueryPharmacistCitizenPrescriptionGet struct {
	CitizenUCN string `query:"citizenUcn"`
}

type RequestPharmacistCitizenFulfillWholePrescription struct {
	Prescriptions []struct {
		Id uuid.UUID `json:"id"`
	} `json:"prescriptions"`
}

//type RequestPharmacistCitizenFulfillMedicamentFromPrescription struct {
//	CitizenId      uuid.UUID `json:"citizen_id"`
//	PrescriptionId uuid.UUID `json:"prescription_id"`
//	MedicamentId   uuid.UUID `json:"medicament_id"`
//}

type RequestPharmacistCitizenFulfillMedicamentFromPrescription struct {
	Prescriptions []struct {
		Id          uuid.UUID `json:"id"`
		Medicaments []struct {
			Id uuid.UUID `json:"id"`
		} `json:"medicaments"`
	} `json:"prescriptions"`
}

type RequestPharmacistBranchAddMedicament struct {
	Medicaments []struct {
		MedicamentId uuid.UUID `json:"id"`
		Quantity     uint      `json:"quantity"`
	} `json:"medicaments"`
}

type ResponsePharmacistCitizenPrescription struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creation_date"`
	StartDate    time.Time `json:"issuedDate"`
	EndDate      time.Time `json:"end_date"`
	Medicaments  []struct {
		Id           uuid.UUID `json:"id"`
		OfficialName string    `json:"officialName"`
		Quantity     uint      `json:"quantity"`
		Fulfilled    bool      `json:"fulfilled"`
	} `json:"medicaments"`
}
