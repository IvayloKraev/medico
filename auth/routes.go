package auth

import (
	"github.com/gofiber/fiber/v2"
)

func SetUpAuthenticationRoutes(router fiber.Router) {
	route := router.Group("/auth")

	controller := newController()

	route.Post("/admin/login", controller.loginAdmin)
}
