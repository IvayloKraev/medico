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

	setupDoctorRoutes(apiRoute)

	setupCitizenRoute(apiRoute)

	pharmacyRoute := apiRoute.Group("/pharmacy")

	setupPharmacyOwnerRoute(pharmacyRoute)
	setupPharmacistsRoute(pharmacyRoute)
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
		"http://localhost:3000",
		"medico.online",
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(allowedOrigins, ","),
		AllowMethods:     strings.Join(allowedMethods, ","),
		AllowHeaders:     strings.Join(allowedHeaders, ","),
		AllowCredentials: true,
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

func setupDoctorRoutes(route fiber.Router) {
	doctor := controllers.NewDoctorController()

	doctorRoute := route.Group("/doctor")
	doctorRoute.Use(doctor.VerifySession)
	doctorRoute.Post("/login", doctor.Login)
	doctorRoute.Post("/logout", doctor.Logout)

	doctorRoute.Get("/citizen/info", doctor.GetCitizenInfo)
	doctorRoute.Get("/citizen/prescription", doctor.GetCitizenPrescriptions)
	doctorRoute.Post("/citizen/prescription", doctor.CreateCitizenPrescription)
}

func setupCitizenRoute(router fiber.Router) {
	citizen := controllers.NewCitizenController()

	citizenRoute := router.Group("/citizen")
	citizenRoute.Use(citizen.VerifySession)
	citizenRoute.Post("/login", citizen.Login)
	citizenRoute.Post("/logout", citizen.Logout)
	citizenRoute.Get("/medicalInfo", citizen.GetMedicalInfo)
	citizenRoute.Get("/personalDoctor", citizen.GetPersonalDoctor)
	citizenRoute.Get("/prescriptions", citizen.Prescription)
	citizenRoute.Get("/available_pharmacies", citizen.AvailablePharmacies)
}

func setupPharmacyOwnerRoute(router fiber.Router) {
	pharmacy := controllers.NewPharmacyOwnerController()

	pharmacyRoute := router.Group("/owner")
	pharmacyRoute.Use(pharmacy.VerifySession)
	pharmacyRoute.Post("/login", pharmacy.Login)
	pharmacyRoute.Post("/logout", pharmacy.Logout)
	pharmacyRoute.Get("/branches", pharmacy.GetAllBranches)
	pharmacyRoute.Get("/pharmacists", pharmacy.GetAllPharmacists)
	pharmacyRoute.Post("/branch/new", pharmacy.NewPharmacyBranch)
	pharmacyRoute.Post("/pharmacist/new", pharmacy.NewPharmacist)
}

func setupPharmacistsRoute(router fiber.Router) {
	pharmacist := controllers.NewPharmacistController()

	pharmacistRoute := router.Group("/pharmacist")
	pharmacistRoute.Use(pharmacist.VerifySession)
	pharmacistRoute.Post("/login", pharmacist.Login)
	pharmacistRoute.Post("/logout", pharmacist.Logout)
	pharmacistRoute.Get("/prescription/get", pharmacist.GetCitizenPrescription)
	pharmacistRoute.Post("/prescription/fulfill", pharmacist.FulfillPrescription)
	pharmacistRoute.Post("/prescription/fulfillMedicament", pharmacist.FulfillMedicamentFromPrescription)
	pharmacistRoute.Post("/branch/addMedicament", pharmacist.AddMedicamentToBranchStorage)
}
