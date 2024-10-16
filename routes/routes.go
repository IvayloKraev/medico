package routes

import (
	"github.com/gofiber/fiber/v2"
	"medico/controllers"
)

func SetUpRoutes(app *fiber.App) {
	authApi := app.Group("/auth")
	authApi.Post("/login", controllers.SignIn)
}
