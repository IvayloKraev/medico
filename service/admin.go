package service

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
