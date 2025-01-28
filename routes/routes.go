package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/session"
	"medico/controllers"
	"strings"
)

func SetUpRoutes(app *fiber.App) {
	apiRoute := app.Group("/api")

	setUpCORS(apiRoute)
	//setUpCSRF(apiRoute)

	setUpCitizenRoute(apiRoute)
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
		"*",
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

	store := session.New(session.Config{
		KeyLookup: "cookie:csrf_medico",
	})

	router.Use(csrf.New(csrf.Config{
		KeyLookup:  "cookie:csrf_medico",
		CookieName: "csrf_medico",
		Session:    store,
		//Extractor:      csrf.CsrfFromCookie("csrf_medico"),
	}))

	router.Get("/csrf-token", func(c *fiber.Ctx) error {
		token := c.Cookies("csrf_medico")
		return c.JSON(fiber.Map{"csrf_medico": token})
	})

}

func setUpCitizenRoute(router fiber.Router) {
	citizen := controllers.NewCitizenController()

	citizenRoute := router.Group("/citizen")
	citizenRoute.Post("/login", citizen.Login)
	citizenRoute.Get("/prescriptions", citizen.Prescription)
	citizenRoute.Get("/available_pharmacies", citizen.AvailablePharmacies)
}
