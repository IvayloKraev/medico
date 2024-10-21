package main

import (
	"github.com/gofiber/fiber/v2"
	"medico/routes"
)

func main() {
	medicoFiber := fiber.New()

	routes.SetUpRoutes(medicoFiber)

	medicoFiber.Listen("0.0.0.0:3000")
}
