package models

type Hospital struct {
	ID      ModelID `gorm:"primaryKey;unique;type:uuid;not null"`
	Name    Text
	Address Text
	Doctors []Doctor `gorm:"foreignKey:HospitalID"`
}

type Doctor struct {
	ID         ModelID `gorm:"primaryKey;unique;type:uuid;not null"`
	FirstName  Text
	SecondName Text
	Surname    Text
	HospitalID ModelID  `gorm:"type:uuid;not null"`
	Hospital   Hospital `gorm:"foreignKey:HospitalID"`
	Uin        Text
}
