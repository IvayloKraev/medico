package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"medico/data"
	"medico/db"
	"medico/dto"
)

func SignIn(c *fiber.Ctx) error {
	signInData := new(dto.CommonUserSignIn)

	if err := c.BodyParser(signInData); err != nil {

		fmt.Println(err)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": true,
			"msg": err.Error(),
		})
	}

	if err := signInData.Validate(); err != nil {

		fmt.Println(err)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": true,
			"msg": err.Error(),
		})
	}

	cmu := data.CommonUserDB{
		ID:           uuid.New(),
		FirstName:    "",
		SecondName:   "",
		LastName:     "",
		Email:        signInData.Email,
		Password:     signInData.Password,
		PasswordSalt: "",
	}

	db.DBConn.Create(&cmu)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
