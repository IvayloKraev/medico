package repo

import (
	"medico/config"
	"medico/models"
)

type AdminRepo interface {
	FindAuthByEmail(email string, adminAuth *models.AdminAuth) error
}

type adminRepo struct {
	repo Repository
}

func NewAdminRepo() AdminRepo {
	databaseConfig := config.LoadDatabaseConfig()
	return &adminRepo{repo: CreateNewRepository(databaseConfig)}
}

func (r *adminRepo) FindAuthByEmail(email string, adminAuth *models.AdminAuth) error {
	return r.repo.First(adminAuth, "email = ?", email).Error
}
