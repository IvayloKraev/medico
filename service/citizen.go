package service

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"medico/models"
	"medico/repo"
	"medico/session"
	"time"
)

type CitizenService interface {
	AuthenticateByEmailAndPassword(email string, password string, citizenAuth *models.CitizenAuth) error
	CreateAuthenticationSession(citizenId uuid.UUID) (uuid.UUID, time.Duration, error)
	GetAuthenticationSession(sessionID uuid.UUID) (uuid.UUID, error)
	DeleteAuthenticationSession(sessionID uuid.UUID) error
	FindAllAvailablePharmacies() error
	ListPrescriptions() error
}

type citizenService struct {
	authSession session.AuthSession
	citizenRepo repo.CitizenRepo
}

func NewCitizenService() CitizenService {
	return &citizenService{
		authSession: session.NewAuthSession("citizen"),
		citizenRepo: repo.NewCitizenRepo(),
	}
}

func (c *citizenService) AuthenticateByEmailAndPassword(email string, password string, citizenAuth *models.CitizenAuth) error {

	if err := c.citizenRepo.FindAuthByEmail(email, citizenAuth); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(citizenAuth.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}

func (c *citizenService) CreateAuthenticationSession(citizenId uuid.UUID) (uuid.UUID, time.Duration, error) {
	return c.authSession.CreateAuthSession(citizenId)
}

func (c *citizenService) GetAuthenticationSession(sessionId uuid.UUID) (uuid.UUID, error) {
	return c.authSession.GetAuthSession(sessionId)
}

func (c *citizenService) DeleteAuthenticationSession(sessionID uuid.UUID) error {
	return c.authSession.DeleteAuthSession(sessionID)
}

func (c *citizenService) FindAllAvailablePharmacies() error {
	return errors.New("not implemented")
}

func (c *citizenService) ListPrescriptions() error {
	return errors.New("not implemented")
}
