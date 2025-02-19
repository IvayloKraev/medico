package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"medico/dto"
	"medico/models"
	"medico/service"
	"regexp"
	"time"
)

type ModeratorController interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	VerifySession(ctx *fiber.Ctx) error

	GetDoctors(ctx *fiber.Ctx) error
	AddDoctor(ctx *fiber.Ctx) error
	DeleteDoctor(ctx *fiber.Ctx) error

	GetMedicaments(ctx *fiber.Ctx) error
	AddMedicament(ctx *fiber.Ctx) error
	DeleteMedicament(ctx *fiber.Ctx) error

	GetPharmacies(ctx *fiber.Ctx) error
	AddPharmacy(ctx *fiber.Ctx) error
	DeletePharmacy(ctx *fiber.Ctx) error

	GetCitizens(ctx *fiber.Ctx) error
	AddCitizen(ctx *fiber.Ctx) error
	DeleteCitizen(ctx *fiber.Ctx) error
}

type moderatorController struct {
	service service.ModeratorService
}

func NewModeratorController() ModeratorController {
	return &moderatorController{service: service.NewModeratorService()}
}

func (m *moderatorController) Login(ctx *fiber.Ctx) error {
	moderatorLogin := new(dto.ModeratorLogin)

	if err := ctx.BodyParser(moderatorLogin); err != nil {
		return err
	}

	moderatorId := uuid.UUID{}
	if err := m.service.AuthenticateWithEmailAndPassword(moderatorLogin.Email.ToString(), moderatorLogin.Password.ToString(), &moderatorId); err != nil {
		return err
	}

	moderator := models.Moderator{}
	if err := m.service.GetModeratorDetails(moderatorId, &moderator); err != nil {
		return err
	}

	session, expiry, err := m.service.CreateAuthenticationSession(moderatorId, moderator.Type)
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Value:   session.String(),
		Expires: time.Now().Add(expiry),
	})

	ctx.Cookie(&fiber.Cookie{
		Name:    "moderator_type",
		Value:   string(moderator.Type),
		Expires: time.Now().Add(expiry),
	})

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

func (m *moderatorController) Logout(ctx *fiber.Ctx) error {
	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	moderatorType := models.ModeratorType(ctx.Cookies("moderator_type", ""))

	if err := m.service.DeleteAuthenticationSession(sessionId, moderatorType); err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Expires: time.Now().Add(-(time.Hour * 2)),
	})

	ctx.Cookie(&fiber.Cookie{
		Name:    "moderator_type",
		Expires: time.Now().Add(-(time.Hour * 2)),
	})

	return ctx.Status(200).JSON(nil)
}

func (m *moderatorController) VerifySession(ctx *fiber.Ctx) error {
	if ctx.Path() == "/api/moderator/login" {
		return ctx.Next()
	}

	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	moderatorType, err := models.ModeratorTypeFromText(ctx.Cookies("moderator_type", ""))
	if err != nil {
		return err
	}

	moderatorId, err := m.service.GetAuthenticationSession(sessionId, moderatorType)
	if err != nil {
		return err
	}

	doctorPathRegex, err := regexp.Compile(`/api/moderator/[a-z]+_doctors?`)
	if err != nil {
		return err
	}

	pharmaPathRegex, err := regexp.Compile(`/api/moderator/[a-z]+_pharmac(ies|y)`)
	if err != nil {
		return err
	}

	citizenPathRegex, err := regexp.Compile(`/api/moderator/[a-z]+_citizens?`)
	if err != nil {
		return err
	}

	medicamentPathRegex, err := regexp.Compile(`/api/moderator/[a-z]+_medicaments?`)
	if err != nil {
		return err
	}

	if !(doctorPathRegex.MatchString(ctx.Path()) && moderatorType == models.DoctorMod) &&
		!(pharmaPathRegex.MatchString(ctx.Path()) && moderatorType == models.PharmacyMod) &&
		!(citizenPathRegex.MatchString(ctx.Path()) && moderatorType == models.CitizenMod) &&
		!(medicamentPathRegex.MatchString(ctx.Path()) && moderatorType == models.MedicamentMod) {
		return errors.New("mismatched role")
	}

	ctx.Locals("moderatorId", moderatorId)

	return ctx.Next()
}

func (m *moderatorController) GetDoctors(ctx *fiber.Ctx) error {
	doctors := new([]dto.ModeratorGetDoctors)

	if err := m.service.FindAllDoctors(doctors); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(doctors)
}

func (m *moderatorController) AddDoctor(ctx *fiber.Ctx) error {
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

func (m *moderatorController) DeleteDoctor(ctx *fiber.Ctx) error {
	doctorId := new(dto.ModeratorDeleteDoctor)

	if err := ctx.BodyParser(&doctorId); err != nil {
		return err
	}

	if err := m.service.DeleteDoctor(doctorId); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

func (m *moderatorController) GetMedicaments(ctx *fiber.Ctx) error {
	medicaments := new([]dto.ModeratorGetMedicaments)

	if err := m.service.FindAllMedicaments(medicaments); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(medicaments)
}

func (m *moderatorController) AddMedicament(ctx *fiber.Ctx) error {
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

func (m *moderatorController) DeleteMedicament(ctx *fiber.Ctx) error {
	medicamentId := new(dto.ModeratorDeleteMedicament)

	if err := ctx.BodyParser(&medicamentId); err != nil {
		return err
	}

	if err := m.service.DeleteMedicament(medicamentId); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

func (m *moderatorController) GetPharmacies(ctx *fiber.Ctx) error {
	pharmacies := new([]dto.ModeratorGetPharmacies)

	if err := m.service.FindAllPharmacies(pharmacies); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(pharmacies)
}

func (m *moderatorController) AddPharmacy(ctx *fiber.Ctx) error {
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

func (m *moderatorController) DeletePharmacy(ctx *fiber.Ctx) error {
	pharmacyId := new(dto.ModeratorDeletePharmacy)

	if err := ctx.BodyParser(&pharmacyId); err != nil {
		return err
	}

	if err := m.service.DeletePharmacy(pharmacyId); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

func (m *moderatorController) GetCitizens(ctx *fiber.Ctx) error {
	citizens := new([]dto.ModeratorGetCitizens)

	if err := m.service.FindAllCitizens(citizens); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(citizens)
}

func (m *moderatorController) AddCitizen(ctx *fiber.Ctx) error {
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

func (m *moderatorController) DeleteCitizen(ctx *fiber.Ctx) error {
	citizenId := new(dto.ModeratorDeleteCitizen)

	if err := ctx.BodyParser(&citizenId); err != nil {
		return err
	}

	if err := m.service.DeleteCitizen(citizenId); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
