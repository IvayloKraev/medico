package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"medico/data"
)

var (
	DBConn *gorm.DB
)

func Connect() {
	fmt.Println("Connecting to database...")
	dsn := "medico:medico@(127.0.0.1:3306)/medico?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	//err = db.Migrator().DropTable(&data.CommonUserDB{})
	//if err != nil {
	//	fmt.Println("Failed to drop table. \n", err)
	//	return
	//}
	//fmt.Println("Successfully drop table. \n", err)

	err = db.AutoMigrate(&data.CommonUserDB{})
	if err != nil {
		fmt.Println("Failed to migrate table. \n", err)
		return
	}
	fmt.Println("Successfully migrate table. \n", err)

	DBConn = db

}
