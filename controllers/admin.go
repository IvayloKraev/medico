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

type AdminController interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	VerifySession(ctx *fiber.Ctx) error
	GetModerators(ctx *fiber.Ctx) error
	AddModerator(ctx *fiber.Ctx) error
	DeleteModerator(ctx *fiber.Ctx) error
}

type adminController struct {
	service service.AdminService
}

func NewAdminController() AdminController {
	return &adminController{service: service.NewAdminService()}
}

func (c *adminController) Login(ctx *fiber.Ctx) error {
	adminLogin := new(dto.AdminLogin)

	if err := ctx.BodyParser(&adminLogin); err != nil {
		return err
	}

	if err := adminLogin.Validate(); err != nil {
		return err
	}

	adminAuth := models.AdminAuth{}

	if err := c.service.AuthenticateByEmailAndPassword(adminLogin.Password.ToString(), adminLogin.Password.ToString(), &adminAuth); err != nil {
		return err
	}

	session, expiry, err := c.service.CreateAuthenticationSession(adminAuth.ID)
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

func (c *adminController) Logout(ctx *fiber.Ctx) error {
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

func (c *adminController) VerifySession(ctx *fiber.Ctx) error {
	if ctx.Path() == "/api/admin/login" {
		return ctx.Next()
	}

	sessionId, err := uuid.Parse(ctx.Cookies("medico_session", uuid.Nil.String()))
	if err != nil {
		return err
	}

	if sessionId == uuid.Nil {
		return errors.New("not logged in")
	}

	adminId, err := c.service.GetAuthenticationSession(sessionId)
	if err != nil {
		return err
	}

	ctx.Locals("adminId", adminId)

	return ctx.Next()
}

func (c *adminController) GetModerators(ctx *fiber.Ctx) error {
	return nil
}

func (c *adminController) AddModerator(ctx *fiber.Ctx) error {
	return nil
}

func (c *adminController) DeleteModerator(ctx *fiber.Ctx) error {
	return nil
}
