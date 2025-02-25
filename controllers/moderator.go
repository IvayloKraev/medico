package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"medico/dto"
	"medico/models"
	"medico/service"
	"time"
)

// DOCTOR

type DoctorModeratorController interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	VerifySession(ctx *fiber.Ctx) error

	GetDoctors(ctx *fiber.Ctx) error
	AddDoctor(ctx *fiber.Ctx) error
	DeleteDoctor(ctx *fiber.Ctx) error
}

type doctorModeratorController struct {
	service service.DoctorModeratorService
}

func NewDoctorModeratorController() DoctorModeratorController {
	return &doctorModeratorController{
		service: service.NewDoctorModeratorService(),
	}
}

func (m *doctorModeratorController) Login(ctx *fiber.Ctx) error {
	moderatorLogin := new(dto.ModeratorLogin)

	if err := ctx.BodyParser(moderatorLogin); err != nil {
		return err
	}

	moderatorId, err := m.service.AuthenticateWithEmailAndPassword(moderatorLogin.Email.ToString(), moderatorLogin.Password.ToString())
	if err != nil {
		return err
	}

	moderator := models.Moderator{}
	if err := m.service.GetModeratorDetails(moderatorId, &moderator); err != nil {
		return err
	}

	session, expiry, err := m.service.CreateAuthenticationSession(moderatorId)
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Value:   session.String(),
		Expires: time.Now().Add(expiry),
	})

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
func (m *doctorModeratorController) Logout(ctx *fiber.Ctx) error {
	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	if err := m.service.DeleteAuthenticationSession(sessionId); err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Expires: time.Now().Add(-(time.Hour * 2)),
	})

	return ctx.Status(200).JSON(nil)
}
func (m *doctorModeratorController) VerifySession(ctx *fiber.Ctx) error {
	if ctx.Path() == "/api/moderator/doctor/login" {
		return ctx.Next()
	}

	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	moderatorId, err := m.service.GetAuthenticationSession(sessionId)
	if err != nil {
		return err
	}

	ctx.Locals("moderatorId", moderatorId)

	return ctx.Next()
}

func (m *doctorModeratorController) GetDoctors(ctx *fiber.Ctx) error {
	doctors := new([]dto.ModeratorGetDoctors)

	if err := m.service.FindAllDoctors(doctors); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(doctors)
}
func (m *doctorModeratorController) AddDoctor(ctx *fiber.Ctx) error {
	newDoctor := new(dto.ModeratorCreateDoctor)

	if err := ctx.BodyParser(&newDoctor); err != nil {
		return err
	}

	err := m.service.CreateDoctor(newDoctor)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(nil)
}
func (m *doctorModeratorController) DeleteDoctor(ctx *fiber.Ctx) error {
	doctorId := new(dto.ModeratorDeleteDoctor)

	if err := ctx.BodyParser(&doctorId); err != nil {
		return err
	}

	if err := m.service.DeleteDoctor(doctorId); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

// PHARMA

type PharmaModeratorController interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	VerifySession(ctx *fiber.Ctx) error

	GetPharmacies(ctx *fiber.Ctx) error
	AddPharmacy(ctx *fiber.Ctx) error
	DeletePharmacy(ctx *fiber.Ctx) error
}

type pharmaModeratorController struct {
	service service.PharmaModeratorService
}

func NewPharmaModeratorController() PharmaModeratorController {
	return &pharmaModeratorController{
		service: service.NewPharmaModeratorService(),
	}
}

func (m *pharmaModeratorController) Login(ctx *fiber.Ctx) error {
	moderatorLogin := new(dto.ModeratorLogin)

	if err := ctx.BodyParser(moderatorLogin); err != nil {
		return err
	}

	moderatorId, err := m.service.AuthenticateWithEmailAndPassword(moderatorLogin.Email.ToString(), moderatorLogin.Password.ToString())
	if err != nil {
		return err
	}

	moderator := models.Moderator{}
	if err := m.service.GetModeratorDetails(moderatorId, &moderator); err != nil {
		return err
	}

	session, expiry, err := m.service.CreateAuthenticationSession(moderatorId)
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Value:   session.String(),
		Expires: time.Now().Add(expiry),
	})

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
func (m *pharmaModeratorController) Logout(ctx *fiber.Ctx) error {
	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	if err := m.service.DeleteAuthenticationSession(sessionId); err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Expires: time.Now().Add(-(time.Hour * 2)),
	})

	return ctx.Status(200).JSON(nil)
}
func (m *pharmaModeratorController) VerifySession(ctx *fiber.Ctx) error {
	if ctx.Path() == "/api/moderator/pharma/login" {
		return ctx.Next()
	}

	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	moderatorId, err := m.service.GetAuthenticationSession(sessionId)
	if err != nil {
		return err
	}

	ctx.Locals("moderatorId", moderatorId)

	return ctx.Next()
}

func (m *pharmaModeratorController) GetPharmacies(ctx *fiber.Ctx) error {
	pharmacies := new([]dto.ModeratorGetPharmacies)

	if err := m.service.FindAllPharmacies(pharmacies); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(pharmacies)
}
func (m *pharmaModeratorController) AddPharmacy(ctx *fiber.Ctx) error {
	newPharmacy := new(dto.ModeratorCreatePharmacy)

	if err := ctx.BodyParser(&newPharmacy); err != nil {
		return err
	}

	err := m.service.CreatePharmacyAndOwner(newPharmacy)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(nil)
}
func (m *pharmaModeratorController) DeletePharmacy(ctx *fiber.Ctx) error {
	pharmacyId := new(dto.ModeratorDeletePharmacy)

	if err := ctx.BodyParser(&pharmacyId); err != nil {
		return err
	}

	if err := m.service.DeletePharmacy(pharmacyId); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

// MEDICAMENT

type MedicamentModeratorController interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	VerifySession(ctx *fiber.Ctx) error

	GetMedicaments(ctx *fiber.Ctx) error
	AddMedicament(ctx *fiber.Ctx) error
	DeleteMedicament(ctx *fiber.Ctx) error
}

type medicamentModeratorController struct {
	service service.MedicamentModeratorService
}

func NewMedicamentModeratorController() MedicamentModeratorController {
	return &medicamentModeratorController{
		service: service.NewMedicamentModeratorService(),
	}
}

func (m *medicamentModeratorController) Login(ctx *fiber.Ctx) error {
	moderatorLogin := new(dto.ModeratorLogin)

	if err := ctx.BodyParser(moderatorLogin); err != nil {
		return err
	}

	moderatorId, err := m.service.AuthenticateWithEmailAndPassword(moderatorLogin.Email.ToString(), moderatorLogin.Password.ToString())
	if err != nil {
		return err
	}

	moderator := models.Moderator{}
	if err := m.service.GetModeratorDetails(moderatorId, &moderator); err != nil {
		return err
	}

	session, expiry, err := m.service.CreateAuthenticationSession(moderatorId)
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Value:   session.String(),
		Expires: time.Now().Add(expiry),
	})

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
func (m *medicamentModeratorController) Logout(ctx *fiber.Ctx) error {
	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	if err := m.service.DeleteAuthenticationSession(sessionId); err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Expires: time.Now().Add(-(time.Hour * 2)),
	})

	return ctx.Status(200).JSON(nil)
}
func (m *medicamentModeratorController) VerifySession(ctx *fiber.Ctx) error {
	if ctx.Path() == "/api/moderator/medicament/login" {
		return ctx.Next()
	}

	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	moderatorId, err := m.service.GetAuthenticationSession(sessionId)
	if err != nil {
		return err
	}

	ctx.Locals("moderatorId", moderatorId)

	return ctx.Next()
}

func (m *medicamentModeratorController) GetMedicaments(ctx *fiber.Ctx) error {
	medicaments := new([]dto.ModeratorGetMedicaments)

	if err := m.service.FindAllMedicaments(medicaments); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(medicaments)
}
func (m *medicamentModeratorController) AddMedicament(ctx *fiber.Ctx) error {
	newMedicament := new(dto.ModeratorCreateMedicament)

	if err := ctx.BodyParser(&newMedicament); err != nil {
		return err
	}

	err := m.service.CreateMedicament(newMedicament)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(nil)
}
func (m *medicamentModeratorController) DeleteMedicament(ctx *fiber.Ctx) error {
	medicamentId := new(dto.ModeratorDeleteMedicament)

	if err := ctx.BodyParser(&medicamentId); err != nil {
		return err
	}

	if err := m.service.DeleteMedicament(medicamentId); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

// CITIZEN

type CitizenModeratorController interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	VerifySession(ctx *fiber.Ctx) error

	GetCitizens(ctx *fiber.Ctx) error
	AddCitizen(ctx *fiber.Ctx) error
	DeleteCitizen(ctx *fiber.Ctx) error
}

type citizenModeratorController struct {
	service service.CitizenModeratorService
}

func NewCitizenModeratorController() CitizenModeratorController {
	return &citizenModeratorController{
		service: service.NewCitizenModeratorService(),
	}
}

func (m *citizenModeratorController) Login(ctx *fiber.Ctx) error {
	moderatorLogin := new(dto.ModeratorLogin)

	if err := ctx.BodyParser(moderatorLogin); err != nil {
		return err
	}

	moderatorId, err := m.service.AuthenticateWithEmailAndPassword(moderatorLogin.Email.ToString(), moderatorLogin.Password.ToString())
	if err != nil {
		return err
	}

	moderator := models.Moderator{}
	if err := m.service.GetModeratorDetails(moderatorId, &moderator); err != nil {
		return err
	}

	session, expiry, err := m.service.CreateAuthenticationSession(moderatorId)
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Value:   session.String(),
		Expires: time.Now().Add(expiry),
	})

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
func (m *citizenModeratorController) Logout(ctx *fiber.Ctx) error {
	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	if err := m.service.DeleteAuthenticationSession(sessionId); err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Expires: time.Now().Add(-(time.Hour * 2)),
	})

	return ctx.Status(200).JSON(nil)
}
func (m *citizenModeratorController) VerifySession(ctx *fiber.Ctx) error {
	if ctx.Path() == "/api/moderator/citizen/login" {
		return ctx.Next()
	}

	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	moderatorId, err := m.service.GetAuthenticationSession(sessionId)
	if err != nil {
		return err
	}

	ctx.Locals("moderatorId", moderatorId)

	return ctx.Next()
}

func (m *citizenModeratorController) GetCitizens(ctx *fiber.Ctx) error {
	citizens := new([]dto.ModeratorGetCitizens)

	if err := m.service.FindAllCitizens(citizens); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(citizens)
}
func (m *citizenModeratorController) AddCitizen(ctx *fiber.Ctx) error {
	newCitizen := new(dto.ModeratorCreateCitizen)

	if err := ctx.BodyParser(&newCitizen); err != nil {
		return err
	}

	err := m.service.CreateCitizen(newCitizen)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(nil)
}
func (m *citizenModeratorController) DeleteCitizen(ctx *fiber.Ctx) error {
	citizenId := new(dto.ModeratorDeleteCitizen)

	if err := ctx.BodyParser(&citizenId); err != nil {
		return err
	}

	if err := m.service.DeleteCitizen(citizenId); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
