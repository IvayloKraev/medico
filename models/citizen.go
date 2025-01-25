package models

type Sex string

const (
	Male   Sex = "M"
	Female Sex = "F"
)

type Province string

const (
	VarnaProvince Province = "varna"
)

type Municipality string

const (
	VarnaMunicipality Municipality = "varna"
)

type City string

const (
	VarnaCity City = "varna"
)

type Citizen struct {
	ID               ModelID `gorm:"primaryKey;unique;type:uuid;not null;"`
	FirstName        Text
	SecondName       Text
	Surname          Text
	Birthday         DateTime
	Sex              Sex  `gorm:"default:'male';type:enum('male','female');not null;"`
	UCN              Text `gorm:"size:10"`
	Height           float32
	Weight           float32
	Email            Text
	PhoneNumber      Text
	AddressID        ModelID        `gorm:"type:uuid;not null"`
	Address          CitizenAddress `gorm:"foreignKey:AddressID;references:ID;"`
	PersonalDoctorID ModelID        `gorm:"type:uuid;not null;"`
	PersonalDoctor   Doctor         `gorm:"foreignKey:PersonalDoctorID;references:ID;"`
	Prescriptions    []Prescription `gorm:"foreignKey:CitizenID;"`
}

type CitizenAddress struct {
	ID                  ModelID      `gorm:"primaryKey;unique;type:uuid;not null"`
	Province            Province     `gorm:"type:enum('varna')not null;"`
	Municipality        Municipality `gorm:"type:enum('varna');not null;"`
	City                City         `gorm:"type:enum('varna');not null;"`
	NeighbourhoodStreet Text
	StreetUnitNumber    uint16
	Entrance            uint8
	Floor               uint8
	Apartment           uint8
}
