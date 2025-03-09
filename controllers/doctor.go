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

type DoctorController interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	VerifySession(ctx *fiber.Ctx) error

	GetCitizenInfo(ctx *fiber.Ctx) error
	GetMedicamentByCommonName(ctx *fiber.Ctx) error
	GetListOfCitizensViaCommonUCN(ctx *fiber.Ctx) error
	GetCitizenPrescriptions(ctx *fiber.Ctx) error
	CreateCitizenPrescription(ctx *fiber.Ctx) error
}

type doctorController struct {
	service service.DoctorService
}

func NewDoctorController() DoctorController {
	return &doctorController{
		service: service.NewDoctorService(),
	}
}

func (d *doctorController) Login(ctx *fiber.Ctx) error {
	doctorLogin := new(dto.RequestDoctorLogin)

	if err := ctx.BodyParser(&doctorLogin); err != nil {
		return err
	}

	doctorAuth := models.DoctorAuth{}

	if err := d.service.AuthenticateByEmailAndPassword(doctorLogin.Email, doctorLogin.Password, &doctorAuth); err != nil {
		return err
	}

	session, expiry, err := d.service.CreateAuthenticationSession(doctorAuth.ID)
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Value:   session.String(),
		Expires: time.Now().Add(expiry),
	})

	return nil
}

func (d *doctorController) Logout(ctx *fiber.Ctx) error {
	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	if err := d.service.DeleteAuthenticationSession(sessionId); err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Expires: time.Now().Add(-(time.Hour * 2)),
	})

	return ctx.Status(200).JSON(nil)
}

func (d *doctorController) VerifySession(ctx *fiber.Ctx) error {

	if ctx.Path() == "/api/doctor/login" {
		return ctx.Next()
	}

	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	doctorId, err := d.service.GetAuthenticationSession(sessionId)
	if err != nil {
		return err
	}

	ctx.Locals("doctorId", doctorId)

	return ctx.Next()
}

func (d *doctorController) GetCitizenInfo(ctx *fiber.Ctx) error {
	citizenUcnDto := new(dto.QueryDoctorGetCitizenInfo)

	if err := ctx.QueryParser(citizenUcnDto); err != nil {
		return err
	}

	citizenInfoDto := new(dto.ResponseDoctorCitizenInfo)

	if err := d.service.GetCitizenInfo(ctx.Locals("doctorId").(uuid.UUID), citizenUcnDto.CitizenUcn, citizenInfoDto); err != nil {
		return err
	}

	return ctx.Status(200).JSON(citizenInfoDto)
}

func (d *doctorController) GetListOfCitizensViaCommonUCN(ctx *fiber.Ctx) error {
	citizenUcnDto := new(dto.QueryDoctorGetCitizenInfo)
	if err := ctx.QueryParser(citizenUcnDto); err != nil {
		return err
	}

	citizensDto := new([]dto.ResponseListOfCitizensViaCommonUCN)

	err := d.service.GetCitizensViaCommonUCN(citizenUcnDto.CitizenUcn, citizensDto)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(citizensDto)
}

func (d *doctorController) GetCitizenPrescriptions(ctx *fiber.Ctx) error {
	citizenIdDto := new(dto.QueryDoctorGetCitizenPrescription)

	if err := ctx.QueryParser(citizenIdDto); err != nil {
		return err
	}

	citizenPrescriptionDto := new([]dto.ResponseDoctorGetCitizenPrescription)

	if err := d.service.GetCitizensPrescriptions(ctx.Locals("doctorId").(uuid.UUID), citizenIdDto.CitizenId, citizenPrescriptionDto); err != nil {
		return err
	}

	return ctx.Status(200).JSON(citizenPrescriptionDto)
}

func (d *doctorController) CreateCitizenPrescription(ctx *fiber.Ctx) error {
	citizenPrescriptionDto := new(dto.RequestDoctorCreatePrescription)

	if err := ctx.BodyParser(&citizenPrescriptionDto); err != nil {
		return err
	}

	if err := d.service.CreatePrescription(ctx.Locals("doctorId").(uuid.UUID), citizenPrescriptionDto); err != nil {
		return err
	}

	return ctx.Status(200).JSON(nil)
}

func (d *doctorController) GetMedicamentByCommonName(ctx *fiber.Ctx) error {
	commonName := new(dto.QueryDoctorGetMedicamentByCommonName)

	if err := ctx.QueryParser(commonName); err != nil {
		return err
	}

	medicamentsDto := new([]dto.ResponseDoctorGetMedicamentPrescription)

	err := d.service.GetMedicamentByCommonName(commonName, medicamentsDto)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(medicamentsDto)
}
