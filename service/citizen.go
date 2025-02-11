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
	AuthenticateByEmailAndPassword(email string, password string) (error, models.CitizenAuth)
	CreateAuthenticateSession(citizenId uuid.UUID) (uuid.UUID, time.Duration, error)
	VerifyAuthenticateSession(sessionID uuid.UUID) (uuid.UUID, error)
	DeleteAuthenticateSession(sessionID uuid.UUID) error
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

func (c *citizenService) AuthenticateByEmailAndPassword(email string, password string) (error, models.CitizenAuth) {

	currentCitizen, err := c.citizenRepo.FindAuthByEmail(email)
	if err != nil {
		return err, models.CitizenAuth{}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(currentCitizen.Password), []byte(password)); err != nil {
		return err, models.CitizenAuth{}
	}

	return nil, currentCitizen
}

func (c *citizenService) CreateAuthenticateSession(citizenId uuid.UUID) (uuid.UUID, time.Duration, error) {
	return c.authSession.CreateAuthSession(citizenId)
}

func (c *citizenService) VerifyAuthenticateSession(sessionId uuid.UUID) (uuid.UUID, error) {
	return c.authSession.GetDataAuthSession(sessionId)
}

func (c *citizenService) DeleteAuthenticateSession(sessionID uuid.UUID) error {
	return c.authSession.DeleteAuthSession(sessionID)
}

func (c *citizenService) FindAllAvailablePharmacies() error {
	return errors.New("not implemented")
}

func (c *citizenService) ListPrescriptions() error {
	return errors.New("not implemented")
}
