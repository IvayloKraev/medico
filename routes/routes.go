package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"medico/controllers"
	"strings"
)

func SetUpRoutes(app *fiber.App) {
	apiRoute := app.Group("/api")

	setUpCORS(apiRoute)

}

func setUpCORS(router fiber.Router) {
	allowedHeaders := []string{
		fiber.HeaderContentType,
		fiber.HeaderAuthorization,
		fiber.HeaderCacheControl,
		fiber.HeaderOrigin,
	}

	allowedMethods := []string{
		fiber.MethodPost,
		fiber.MethodPut,
		fiber.MethodGet,
		fiber.MethodDelete,
		fiber.MethodOptions,
	}

	allowedOrigins := []string{
		"http://localhost:3000",
	}

	router.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins:     strings.Join(allowedOrigins, ","),
		AllowMethods:     strings.Join(allowedMethods, ","),
		AllowHeaders:     strings.Join(allowedHeaders, ","),
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))
}

func setUpCSRF(router fiber.Router) {
	router.Use(csrf.New(csrf.Config{
		CookieName:     "csrf_medico",
		Expiration:     3,
		SingleUseToken: true,
	}))
}

func setUpCitizenRoute(router fiber.Router) {
	citizen := controllers.NewCitizenController()

	citizenRoute := router.Group("/citizen")
	citizenRoute.Post("/login", citizen.Login)
	citizenRoute.Get("/prescriptions", citizen.Prescription)
	citizenRoute.Get("/available_pharmacies", citizen.AvailablePharmacies)
}
