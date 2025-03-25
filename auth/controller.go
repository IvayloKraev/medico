package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type controllerQueries interface {
	verifyAdminSession(sessionId uuid.UUID) error
}

type controllerMutations interface {
	loginAdmin(ctx *fiber.Ctx) error
}

type controller struct {
	service *service
}

var (
	_ controllerQueries   = (*controller)(nil)
	_ controllerMutations = (*controller)(nil)
)

func newController() *controller {
	return &controller{
		service: newService(),
	}
}

func (c controller) verifyAdminSession(sessionId uuid.UUID) error {
	panic("implement me")
}

func (c controller) loginAdmin(*fiber.Ctx) error {
	return fiber.ErrNotImplemented
}
