package dto

type CitizenLogin struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
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
