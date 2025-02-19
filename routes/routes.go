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

func SetupRoutes(app *fiber.App) {
	apiRoute := app.Group("/api")

	setupCORS(apiRoute)
	setupCSRF(apiRoute)

	setupAdminRoutes(apiRoute)

	moderatorRoute := apiRoute.Group("/moderator")

	setupDoctorModeratorRoutes(moderatorRoute)
	setupPharmaModeratorRoutes(moderatorRoute)
	setupMedicamentModeratorRoutes(moderatorRoute)
	setupCitizenModeratorRoutes(moderatorRoute)

	setupCitizenRoute(apiRoute)
}

func setupCORS(router fiber.Router) {
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

func setupCSRF(router fiber.Router) {
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

func setupAdminRoutes(router fiber.Router) {
	admin := controllers.NewAdminController()

	adminRoute := router.Group("/admin")
	adminRoute.Use(admin.VerifySession)
	adminRoute.Post("/login", admin.Login)
	adminRoute.Post("/logout", admin.Logout)
	adminRoute.Get("/moderator/get", admin.GetModerators)
	adminRoute.Post("/moderator/create", admin.AddModerator)
	adminRoute.Delete("/moderator/delete", admin.DeleteModerator)
}

func setupDoctorModeratorRoutes(moderatorRoute fiber.Router) {
	doctorModerator := controllers.NewDoctorModeratorController()

	doctorModeratorRoute := moderatorRoute.Group("/doctor")
	doctorModeratorRoute.Use(doctorModerator.VerifySession)
	doctorModeratorRoute.Post("/login", doctorModerator.Login)
	doctorModeratorRoute.Post("/logout", doctorModerator.Logout)

	doctorModeratorRoute.Get("/get", doctorModerator.GetDoctors)
	doctorModeratorRoute.Post("/create", doctorModerator.AddDoctor)
	doctorModeratorRoute.Delete("/delete", doctorModerator.DeleteDoctor)
}

func setupPharmaModeratorRoutes(moderatorRoute fiber.Router) {
	pharmaModerator := controllers.NewPharmaModeratorController()

	pharmaModeratorRoute := moderatorRoute.Group("/pharma")
	pharmaModeratorRoute.Use(pharmaModerator.VerifySession)
	pharmaModeratorRoute.Post("/login", pharmaModerator.Login)
	pharmaModeratorRoute.Post("/logout", pharmaModerator.Logout)

	pharmaModeratorRoute.Get("/get", pharmaModerator.GetPharmacies)
	pharmaModeratorRoute.Post("/create", pharmaModerator.AddPharmacy)
	pharmaModeratorRoute.Delete("/delete", pharmaModerator.DeletePharmacy)
}

func setupMedicamentModeratorRoutes(moderatorRoute fiber.Router) {
	medicamentModerator := controllers.NewMedicamentModeratorController()

	medicamentModeratorRoute := moderatorRoute.Group("/medicament")
	medicamentModeratorRoute.Use(medicamentModerator.VerifySession)
	medicamentModeratorRoute.Post("/login", medicamentModerator.Login)
	medicamentModeratorRoute.Post("/logout", medicamentModerator.Logout)

	medicamentModeratorRoute.Get("/get", medicamentModerator.GetMedicaments)
	medicamentModeratorRoute.Post("/create", medicamentModerator.AddMedicament)
	medicamentModeratorRoute.Delete("/delete", medicamentModerator.DeleteMedicament)
}

func setupCitizenModeratorRoutes(moderatorRoute fiber.Router) {
	citizenModerator := controllers.NewCitizenModeratorController()

	citizenModeratorRoute := moderatorRoute.Group("/citizen")
	citizenModeratorRoute.Use(citizenModerator.VerifySession)
	citizenModeratorRoute.Post("/login", citizenModerator.Login)
	citizenModeratorRoute.Post("/logout", citizenModerator.Logout)

	citizenModeratorRoute.Get("/get", citizenModerator.GetCitizens)
	citizenModeratorRoute.Post("/create", citizenModerator.AddCitizen)
	citizenModeratorRoute.Delete("/delete", citizenModerator.DeleteCitizen)
}

func setupCitizenRoute(router fiber.Router) {
	citizen := controllers.NewCitizenController()

	citizenRoute := router.Group("/citizen")
	citizenRoute.Use(citizen.VerifySession)
	citizenRoute.Post("/login", citizen.Login)
	citizenRoute.Post("/logout", citizen.Logout)
	citizenRoute.Get("/prescriptions", citizen.Prescription)
	citizenRoute.Get("/available_pharmacies", citizen.AvailablePharmacies)
}
