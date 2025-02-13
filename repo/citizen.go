package repo

import (
	"medico/config"
	"medico/models"
)

type CitizenRepo interface {
	FindAuthByEmail(email string, citizenAuth *models.CitizenAuth) error
}

type citizenRepo struct {
	repo Repository
}

func NewCitizenRepo() CitizenRepo {
	databaseConfig := config.LoadDatabaseConfig()
	return &citizenRepo{repo: CreateNewRepository(databaseConfig)}
}

func (r *citizenRepo) FindAuthByEmail(email string, citizenAuth *models.CitizenAuth) error {
	return r.repo.First(citizenAuth, "email = ?", email).Error
}
