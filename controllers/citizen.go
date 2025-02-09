package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"medico/dto"
	"medico/service"
	"time"
)

type CitizenController interface {
	Login(c *fiber.Ctx) error
	Prescription(c *fiber.Ctx) error
	AvailablePharmacies(c *fiber.Ctx) error
}

type citizenController struct {
	Path    string
	service service.CitizenService
}

// NewCitizenController is constructor for CitizenController
func NewCitizenController() CitizenController {
	return &citizenController{Path: "/citizen", service: service.NewCitizenService()}
}

func (c *citizenController) Login(ctx *fiber.Ctx) error {
	loginData := new(dto.CitizenLogin)

	if err := ctx.BodyParser(loginData); err != nil {
		return err
	}

	if err := loginData.Validate(); err != nil {
		return err
	}

	err, m := c.service.AuthenticateByEmailAndPassword(loginData.Email.ToString(), loginData.Password.ToString())
	if err != nil {
		return err
	}

	session, expiry, err := c.service.CreateAuthenticateSession(m.ID)

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

func (c *citizenController) Prescription(ctx *fiber.Ctx) error {
	return errors.New("not implemented")
}

func (c *citizenController) AvailablePharmacies(ctx *fiber.Ctx) error {
	return errors.New("not implemented")
}
