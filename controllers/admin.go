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
	adminLogin := new(dto.RequestAdminLogin)

	if err := ctx.BodyParser(adminLogin); err != nil {
		return err
	}

	adminAuth := models.AdminAuth{}

	if err := c.service.AuthenticateByEmailAndPassword(adminLogin.Email, adminLogin.Password, &adminAuth); err != nil {
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
	dtoModerators := new([]dto.ResponseAdminGetModerator)

	if err := c.service.GetModerators(dtoModerators); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dtoModerators)
}

func (c *adminController) AddModerator(ctx *fiber.Ctx) error {
	newModerator := new(dto.RequestAdminCreateModerator)

	if err := ctx.BodyParser(newModerator); err != nil {
		return err
	}

	err := c.service.CreateModerator(newModerator)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(nil)
}

func (c *adminController) DeleteModerator(ctx *fiber.Ctx) error {
	moderatorId := new(dto.QueryAdminDeleteModerator)

	if err := ctx.QueryParser(moderatorId); err != nil {
		return err
	}

	if err := c.service.DeleteModerator(moderatorId.ModeratorId); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
