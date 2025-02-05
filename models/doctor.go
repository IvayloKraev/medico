package models

import "github.com/google/uuid"

type Hospital struct {
	ID      uuid.UUID `gorm:"primaryKey;unique;type:uuid;not null"`
	Name    Text
	Address Text
	Doctors []Doctor `gorm:"foreignKey:HospitalID"`
}

type Doctor struct {
	ID         uuid.UUID `gorm:"primaryKey;unique;type:uuid;not null"`
	FirstName  Text
	SecondName Text
	Surname    Text
	HospitalID uuid.UUID `gorm:"type:uuid;not null"`
	Hospital   Hospital  `gorm:"foreignKey:HospitalID"`
	Uin        Text
}
