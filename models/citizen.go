package models

type Citizen struct {
	ID               TModelID `gorm:"primaryKey;unique;type:uuid;not null"`
	FirstName        TShortText
	SecondName       TShortText
	Surname          TShortText
	Age              int
	Height           float32
	Weight           float32
	Email            TShortText
	Address          TShortText
	City             TShortText
	PersonalDoctorID TModelID       `gorm:"type:uuid;not null"`
	PersonalDoctor   Doctor         `gorm:"foreignKey:PersonalDoctorID;references:ID"`
	Prescriptions    []Prescription `gorm:"foreignKey:CitizenID"`
}
