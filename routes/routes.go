package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/storage/redis/v3"
	"medico/config"
	"medico/controllers"
	"strings"
)

func SetUpRoutes(app *fiber.App) {
	apiRoute := app.Group("/api")

	setUpCORS(apiRoute)
	setUpCSRF(apiRoute)

	setUpCitizenRoute(apiRoute)
	setUpAdminRoutes(apiRoute)
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
	csrfConfig := config.LoadCSRFTokenConfig()

	router.Use(csrf.New(csrf.Config{
		CookieName: csrfConfig.CookieName,
		Storage: redis.New(redis.Config{
			Host:     csrfConfig.Host,
			Port:     csrfConfig.Port,
			Username: csrfConfig.Username,
			Reset:    csrfConfig.Reset,
			Database: csrfConfig.Database,
		}),
		Extractor:      csrf.CsrfFromCookie(csrfConfig.CookieName),
		SingleUseToken: csrfConfig.SingleUseToken,
		Expiration:     csrfConfig.Expiration,
	}))

	router.Get("/csrf-token", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(nil)
	})
}

func setUpCitizenRoute(router fiber.Router) {
	citizen := controllers.NewCitizenController()

	citizenRoute := router.Group("/citizen")
	citizenRoute.Use(citizen.VerifySession)
	citizenRoute.Post("/login", citizen.Login)
	citizenRoute.Post("/logout", citizen.Logout)
	citizenRoute.Get("/prescriptions", citizen.Prescription)
	citizenRoute.Get("/available_pharmacies", citizen.AvailablePharmacies)
}

func setUpAdminRoutes(router fiber.Router) {
	admin := controllers.NewAdminController()

	adminRoute := router.Group("/admin")
	adminRoute.Use(admin.VerifySession)
	adminRoute.Post("/login", admin.Login)
	adminRoute.Post("/logout", admin.Logout)
}
