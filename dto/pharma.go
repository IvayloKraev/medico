package dto

import (
	"github.com/google/uuid"
	"time"
)

type PharmacyOwnerAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PharmacyOwnerBranches struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type PharmacyOwnerNewBranch struct {
	Name      string  `json:"name"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type PharmacyOwnerPharmacist struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type PharmacyOwnerNewPharmacist struct {
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	WorkingBranch uuid.UUID `json:"working_branch"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
}

type PharmacistAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PharmacistCitizenPrescriptionGet struct {
	CitizenUCN string `json:"citizen_ucn"`
}

type PharmacistCitizenFulfillWholePrescription struct {
	CitizenId      uuid.UUID `json:"citizen_id"`
	BranchId       uuid.UUID `json:"branch_id"`
	PrescriptionId uuid.UUID `json:"prescription_id"`
}

type PharmacistCitizenFulfillMedicamentFromPrescription struct {
	CitizenId      uuid.UUID `json:"citizen_ucn"`
	PrescriptionId uuid.UUID `json:"prescription_id"`
	MedicamentId   uuid.UUID `json:"medicament_ucn"`
}

type PharmacistBranchAddMedicament struct {
	BranchId    uuid.UUID `json:"branch_id"`
	Medicaments []struct {
		BranchId       uuid.UUID `json:"branch_id"`
		MedicamentName string    `json:"medicament_name"`
		Quantity       uint      `json:"quantity"`
	} `json:"medicaments"`
}

type PharmacistCitizenPrescription struct {
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
