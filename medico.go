package main

import (
	"github.com/gofiber/fiber/v2"
	"medico/config"
	"medico/db"
	"medico/routes"
)

func main() {
	databaseConfig := config.LoadDatabaseConfig()

	mainRepository := db.CreateNewRepository("main", databaseConfig)

	if databaseConfig.Migration {
		db.Migrate(mainRepository)
	}

	medicoFiber := fiber.New()

	routes.SetUpRoutes(medicoFiber)

	_ = medicoFiber.Listen(":8080")
}
