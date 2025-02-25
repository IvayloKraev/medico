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

type CitizenController interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	VerifySession(ctx *fiber.Ctx) error
	GetMedicalInfo(ctx *fiber.Ctx) error
	Prescription(ctx *fiber.Ctx) error
	AvailablePharmacies(ctx *fiber.Ctx) error
}

type citizenController struct {
	service service.CitizenService
}

// NewCitizenController is constructor for CitizenController
func NewCitizenController() CitizenController {
	return &citizenController{service: service.NewCitizenService()}
}

func (c *citizenController) Login(ctx *fiber.Ctx) error {
	loginData := new(dto.CitizenLogin)

	if err := ctx.BodyParser(loginData); err != nil {
		return err
	}

	if err := loginData.Validate(); err != nil {
		return err
	}

	citizenAuth := models.CitizenAuth{}

	if err := c.service.AuthenticateByEmailAndPassword(loginData.Email.ToString(), loginData.Password.ToString(), &citizenAuth); err != nil {
		return err
	}

	session, expiry, err := c.service.CreateAuthenticationSession(citizenAuth.ID)

	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Value:   session.String(),
		Expires: time.Now().Add(expiry),
	})

	return ctx.Status(200).JSON(nil)
}

func (c *citizenController) Logout(ctx *fiber.Ctx) error {
	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	if err := c.service.DeleteAuthenticationSession(sessionId); err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "medico_session",
		Expires: time.Now().Add(-(time.Hour * 2)),
	})

	return ctx.Status(200).JSON(nil)
}

func (c *citizenController) VerifySession(ctx *fiber.Ctx) error {
	if ctx.Path() == "/api/citizen/login" {
		return ctx.Next()
	}

	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	userId, err := c.service.GetAuthenticationSession(sessionId)
	if err != nil {
		return err
	}

	ctx.Locals("citizenId", userId)

	return ctx.Next()
}

func (c *citizenController) GetMedicalInfo(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (c *citizenController) Prescription(ctx *fiber.Ctx) error {
	return errors.New("not implemented")
}

func (c *citizenController) AvailablePharmacies(ctx *fiber.Ctx) error {
	return errors.New("not implemented")
}
