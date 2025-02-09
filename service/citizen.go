package service

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"medico/config"
	"medico/db"
	"medico/models"
	"medico/session"
	"time"
)

type CitizenService interface {
	AuthenticateByEmailAndPassword(email string, password string) (error, models.CitizenAuth)
	CreateAuthenticateSession(citizenId uuid.UUID) (uuid.UUID, time.Duration, error)
	VerifyAuthenticateSession(sessionID uuid.UUID) (uuid.UUID, error)
	FindAllAvailablePharmacies() error
	ListPrescriptions() error
}

type citizenService struct {
	authSession    session.AuthSession
	authRepository db.Repository
}

func NewCitizenService() CitizenService {
	return &citizenService{
		authSession:    session.NewAuthSession("citizen"),
		authRepository: db.CreateNewRepository("Citizen", config.LoadDatabaseConfig()),
	}
}

func (c *citizenService) AuthenticateByEmailAndPassword(email string, password string) (error, models.CitizenAuth) {
	var currentCitizen models.CitizenAuth

	if err := c.authRepository.Where("email = ?", email).First(&currentCitizen).Error; err != nil {
		return err, currentCitizen
	}

	err := bcrypt.CompareHashAndPassword([]byte(currentCitizen.Password), []byte(password))

	if err != nil {
		return err, currentCitizen
	}

	return nil, currentCitizen
}

func (c *citizenService) CreateAuthenticateSession(citizenId uuid.UUID) (uuid.UUID, time.Duration, error) {
	return c.authSession.CreateAuthSession(citizenId)
}

func (c *citizenService) VerifyAuthenticateSession(sessionId uuid.UUID) (uuid.UUID, error) {
	return c.authSession.GetDataAuthSession(sessionId)
}

func (c *citizenService) FindAllAvailablePharmacies() error {
	return errors.New("not implemented")
}

func (c *citizenService) ListPrescriptions() error {
	return errors.New("not implemented")
}
