package main

import (
	"github.com/gofiber/fiber/v2"
	"medico/routes"
	"medico/utils"
)

func init() {
	utils.LoadDatabaseConfig()
	utils.LoadCSRFConfig()
	utils.LoadCSRFConfig()
	utils.LoadHashingCost()
}

func main() {

	medicoFiber := fiber.New()

	routes.SetupRoutes(medicoFiber)

	_ = medicoFiber.Listen(":8080")
}
