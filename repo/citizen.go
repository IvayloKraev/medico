package repo

import (
	"medico/config"
	"medico/models"
)

type CitizenRepo interface {
	FindAuthByEmail(email string) (models.CitizenAuth, error)
}

type citizenRepo struct {
	repo Repository
}

func NewCitizenRepo() CitizenRepo {
	databaseConfig := config.LoadDatabaseConfig()
	return &citizenRepo{repo: CreateNewRepository(databaseConfig)}
}

func (c *citizenRepo) FindAuthByEmail(email string) (models.CitizenAuth, error) {
	citizenAuth := models.CitizenAuth{}
	err := c.repo.First(&citizenAuth, "email = ?", email).Error
	return citizenAuth, err
}
