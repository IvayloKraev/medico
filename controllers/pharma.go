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

type PharmacyOwnerController interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	VerifySession(ctx *fiber.Ctx) error

	GetAllBranches(ctx *fiber.Ctx) error
	GetAllPharmacists(ctx *fiber.Ctx) error
	NewPharmacyBranch(ctx *fiber.Ctx) error
	NewPharmacist(ctx *fiber.Ctx) error
}

type pharmacyOwnerController struct {
	service service.PharmacyOwnerService
}

func NewPharmacyOwnerController() PharmacyOwnerController {
	return &pharmacyOwnerController{
		service: service.NewPharmacyOwnerService(),
	}
}

func (c *pharmacyOwnerController) Login(ctx *fiber.Ctx) error {
	pharmacyOwnerLogin := new(dto.RequestPharmacyOwnerAuth)

	if err := ctx.BodyParser(&pharmacyOwnerLogin); err != nil {
		return err
	}

	pharmacyOwnerAuth := models.PharmacyOwnerAuth{}

	if err := c.service.AuthenticateByEmailAndPassword(pharmacyOwnerLogin.Email, pharmacyOwnerLogin.Password, &pharmacyOwnerAuth); err != nil {
		return err
	}

	session, expiry, err := c.service.CreateAuthenticationSession(pharmacyOwnerAuth.ID)
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

func (c *pharmacyOwnerController) Logout(ctx *fiber.Ctx) error {
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

func (c *pharmacyOwnerController) VerifySession(ctx *fiber.Ctx) error {
	if ctx.Path() == "/api/pharma/owner/login" {
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

	ctx.Locals("pharmacyOwnerId", adminId)

	return ctx.Next()
}

func (c *pharmacyOwnerController) GetAllBranches(ctx *fiber.Ctx) error {
	branches := new([]dto.ResponsePharmacyOwnerBranches)

	if err := c.service.GetAllBranches(ctx.Locals("pharmacyOwnerId").(uuid.UUID), branches); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(*branches)
}

func (c *pharmacyOwnerController) GetAllPharmacists(ctx *fiber.Ctx) error {
	pharmacists := new([]dto.ResponsePharmacyOwnerPharmacist)

	if err := c.service.GetAllPharmacists(ctx.Locals("pharmacyOwnerId").(uuid.UUID), pharmacists); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(*pharmacists)
}

func (c *pharmacyOwnerController) NewPharmacyBranch(ctx *fiber.Ctx) error {
	newBranch := new(dto.RequestPharmacyOwnerNewBranch)

	if err := ctx.BodyParser(&newBranch); err != nil {
		return err
	}

	err := c.service.NewPharmacyBranch(ctx.Locals("pharmacyOwnerId").(uuid.UUID), newBranch)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

func (c *pharmacyOwnerController) NewPharmacist(ctx *fiber.Ctx) error {
	newPharmacist := new(dto.RequestPharmacyOwnerNewPharmacist)
	if err := ctx.BodyParser(&newPharmacist); err != nil {
		return err
	}

	if err := c.service.NewPharmacist(ctx.Locals("pharmacyOwnerId").(uuid.UUID), newPharmacist); err != nil {
		return err
	}

	return nil
}

type PharmacistController interface {
	Login(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
	VerifySession(ctx *fiber.Ctx) error

	GetCitizenPrescription(ctx *fiber.Ctx) error
	FulfillPrescription(ctx *fiber.Ctx) error
	FulfillMedicamentFromPrescription(ctx *fiber.Ctx) error

	AddMedicamentToBranchStorage(ctx *fiber.Ctx) error
}

type pharmacistController struct {
	service service.PharmacistService
}

func NewPharmacistController() PharmacistController {
	return &pharmacistController{
		service: service.NewPharmacistService(),
	}
}

func (c *pharmacistController) Login(ctx *fiber.Ctx) error {
	pharmacistLogin := new(dto.RequestPharmacistAuth)

	if err := ctx.BodyParser(&pharmacistLogin); err != nil {
		return err
	}

	pharmacistAuth := models.PharmacistAuth{}

	if err := c.service.AuthenticateByEmailAndPassword(pharmacistAuth.Email, pharmacistAuth.Password, &pharmacistAuth); err != nil {
		return err
	}

	session, expiry, err := c.service.CreateAuthenticationSession(pharmacistAuth.ID)
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

func (c *pharmacistController) Logout(ctx *fiber.Ctx) error {
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

func (c *pharmacistController) VerifySession(ctx *fiber.Ctx) error {
	if ctx.Path() == "/api/pharma/pharmacist/login" {
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

	ctx.Locals("pharmacistId", adminId)

	return ctx.Next()
}

func (c *pharmacistController) GetCitizenPrescription(ctx *fiber.Ctx) error {
	citizenUcn := new(dto.PharmacistCitizenPrescriptionGet)
	if err := ctx.BodyParser(&citizenUcn); err != nil {
		return err
	}

	data := new([]dto.ResponsePharmacistCitizenPrescription)

	if err := c.service.GetCitizensActivePrescriptions(citizenUcn, data); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(*data)
}

func (c *pharmacistController) FulfillPrescription(ctx *fiber.Ctx) error {
	input := new(dto.RequestPharmacistCitizenFulfillWholePrescription)

	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	if err := c.service.FulfillWholePrescription(input); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

func (c *pharmacistController) FulfillMedicamentFromPrescription(ctx *fiber.Ctx) error {
	input := new(dto.RequestPharmacistCitizenFulfillMedicamentFromPrescription)

	if err := ctx.BodyParser(&input); err != nil {
		return err
	}

	if err := c.service.FulfillMedicamentFromPrescription(input); err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(nil)
}

func (c *pharmacistController) AddMedicamentToBranchStorage(ctx *fiber.Ctx) error {
	input := new(dto.RequestPharmacistBranchAddMedicament)

	if err := ctx.BodyParser(&input); err != nil {
		return err
	}
	if err := c.service.AddMedicamentToBranchStorage(input); err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(nil)
}
