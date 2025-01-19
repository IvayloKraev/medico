package models

type Hospital struct {
	ID      TModelID `gorm:"primaryKey;unique;type:uuid;not null"`
	Name    TShortText
	Address TShortText
	Doctors []Doctor `gorm:"foreignKey:HospitalID"`
}

type Doctor struct {
	ID         TModelID `gorm:"primaryKey;unique;type:uuid;not null"`
	FirstName  TShortText
	SecondName TShortText
	Surname    TShortText
	HospitalID TModelID `gorm:"type:uuid;not null"`
	Hospital   Hospital `gorm:"foreignKey:HospitalID"`
	Uin        TShortText
}
