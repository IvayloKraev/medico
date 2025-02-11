package dto

import "errors"

type CitizenLogin struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
}

func (l CitizenLogin) Validate() error {
	return errors.Join(l.Email.Validate(), l.Password.Validate())
}

type CitizenAvailablePharmacy struct {
	Latitude  float32
	Longitude float32
}

type CitizenPrescription struct {
	Doctor      string
	Medicaments []struct {
		Name string
		Unit uint8
	}
}
