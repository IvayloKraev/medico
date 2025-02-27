package service

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"medico/common"
	"medico/dto"
	"medico/models"
	"medico/repo"
	"medico/session"
	"time"
)

type AdminService interface {
	AuthenticateByEmailAndPassword(email string, password string, adminAuth *models.AdminAuth) error
	CreateAuthenticationSession(adminId uuid.UUID) (uuid.UUID, time.Duration, error)
	GetAuthenticationSession(sessionId uuid.UUID) (uuid.UUID, error)
	DeleteAuthenticationSession(sessionId uuid.UUID) error
	CreateModerator(createModerator *dto.RequestAdminCreateModerator) error
	DeleteModerator(moderatorId uuid.UUID) error
	GetModerators(dtoModerators *[]dto.ResponseAdminGetModerator) error
}

type adminService struct {
	authSession session.AuthSession
	repo        repo.AdminRepo
}

func NewAdminService() AdminService {
	return &adminService{
		authSession: session.NewAuthSession("admin"),
		repo:        repo.NewAdminRepo(),
	}
}

func (s *adminService) AuthenticateByEmailAndPassword(email string, password string, adminAuth *models.AdminAuth) error {
	if err := s.repo.FindAuthByEmail(email, adminAuth); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(adminAuth.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}

func (s *adminService) CreateAuthenticationSession(adminId uuid.UUID) (uuid.UUID, time.Duration, error) {
	return s.authSession.CreateAuthSession(adminId)
}

func (s *adminService) GetAuthenticationSession(sessionId uuid.UUID) (uuid.UUID, error) {
	return s.authSession.GetAuthSession(sessionId)
}

func (s *adminService) DeleteAuthenticationSession(sessionId uuid.UUID) error {
	return s.authSession.DeleteAuthSession(sessionId)
}

func (s *adminService) CreateModerator(createModerator *dto.RequestAdminCreateModerator) error {
	password, err := bcrypt.GenerateFromPassword([]byte(createModerator.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newModeratorAuth := models.ModeratorAuth{
		ID:       uuid.New(),
		Email:    createModerator.Email,
		Password: string(password),
		Moderator: models.Moderator{
			FirstName:  createModerator.FirstName,
			SecondName: createModerator.SecondName,
			LastName:   createModerator.LastName,
			Email:      createModerator.Email,
			Type:       common.ModeratorType(createModerator.Type),
		},
	}

	if err := s.repo.CreateModerator(&newModeratorAuth); err != nil {
		return err
	}

	return nil
}

func (s *adminService) DeleteModerator(moderatorId uuid.UUID) error {
	return s.repo.DeleteModerator(moderatorId)
}

func (s *adminService) GetModerators(dtoModerators *[]dto.ResponseAdminGetModerator) error {
	var moderators []models.Moderator

	if err := s.repo.FindAllModerators(&moderators); err != nil {
		return err
	}

	*dtoModerators = make([]dto.ResponseAdminGetModerator, len(moderators))

	for i, mod := range moderators {
		(*dtoModerators)[i] = dto.ResponseAdminGetModerator{
			ID:         mod.ID,
			FirstName:  mod.FirstName,
			SecondName: mod.SecondName,
			LastName:   mod.LastName,
			Email:      mod.Email,
			Type:       mod.Type,
		}
	}

	return nil
}
