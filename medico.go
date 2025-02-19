package main

import (
	"github.com/gofiber/fiber/v2"
	"medico/config"
	"medico/repo"
	"medico/routes"
)

func main() {
	migrationConfig := config.LoadMigrationConfig()

	if migrationConfig.Migration {
		migrator := repo.NewMigratorRepo()
		err := migrator.MigrateAll()
		if err != nil {
			panic(err)
		}
	}

	medicoFiber := fiber.New()

	routes.SetupRoutes(medicoFiber)

	_ = medicoFiber.Listen(":8080")
}
