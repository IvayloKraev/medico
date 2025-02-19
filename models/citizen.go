package models

import (
	"github.com/google/uuid"
	"time"
)

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

type CitizenAuth struct {
	ID       uuid.UUID `gorm:"primary_key;unique;type:uuid;not null;"`
	Email    string    `gorm:"type:text;not null"`
	Password string    `gorm:"type:text;not null"`
	Citizen  Citizen   `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE;"`
}

type Citizen struct {
	ID         uuid.UUID `gorm:"primaryKey;unique;type:uuid;not null;"`
	FirstName  string
	SecondName string
	LastName   string
	Birthday   time.Time
	Sex        Sex    `gorm:"default:'male';type:enum('male','female');not null;"`
	UCN        string `gorm:"size:10"`
	//Height           float32
	//Weight           float32
	Email       string
	PhoneNumber string
	//AddressID        uuid.UUID      `gorm:"type:uuid;not null"`
	//Address          CitizenAddress `gorm:"foreignKey:AddressID;references:ID;"`
	//PersonalDoctorID uuid.UUID      `gorm:"type:uuid;not null;"`
	//PersonalDoctor   Doctor         `gorm:"foreignKey:PersonalDoctorID;references:ID;"`
	Prescriptions []Prescription `gorm:"foreignKey:CitizenID;"`
}

type CitizenAddress struct {
	ID                  uuid.UUID    `gorm:"primaryKey;unique;type:uuid;not null"`
	Province            Province     `gorm:"type:enum('varna')not null;"`
	Municipality        Municipality `gorm:"type:enum('varna');not null;"`
	City                City         `gorm:"type:enum('varna');not null;"`
	NeighbourhoodStreet string
	StreetUnitNumber    uint16
	Entrance            uint8
	Floor               uint8
	Apartment           uint8
}
