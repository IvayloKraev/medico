package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"medico/models"
)

func SignIn(c *fiber.Ctx) error {
	signInData := new(models.SignIn)

	if err := c.BodyParser(signInData); err != nil {

		fmt.Println(err)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": true,
			"msg": err.Error(),
		})
	}

	fmt.Println(signInData)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
