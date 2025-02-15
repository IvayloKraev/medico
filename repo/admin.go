package repo

import (
	"github.com/google/uuid"
	"medico/config"
	"medico/models"
)

type AdminRepo interface {
	FindAuthByEmail(email string, adminAuth *models.AdminAuth) error
	CreateModerator(moderatorAuth *models.ModeratorAuth) error
	DeleteModerator(moderatorId uuid.UUID) error
	FindAllModerators(moderators *[]models.Moderator) error
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

func (r *adminRepo) CreateModerator(moderatorAuth *models.ModeratorAuth) error {
	return r.repo.Create(moderatorAuth).Error
}

func (r *adminRepo) DeleteModerator(moderatorId uuid.UUID) error {
	return r.repo.Where("id = ?", moderatorId.String()).Delete(models.ModeratorAuth{}).Error
}

func (r *adminRepo) FindAllModerators(moderators *[]models.Moderator) error {
	return r.repo.Find(moderators).Error
}
