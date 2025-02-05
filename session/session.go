package session

import (
	"errors"
	"github.com/gofiber/storage/redis/v3"
	"github.com/google/uuid"
	"medico/config"
	"medico/models"
	"time"
)

type AuthSession interface {
	VerifyAuthSession(sessionId uuid.UUID) error
	CreateAuthSession(citizen models.CitizenAuth) (uuid.UUID, time.Duration, error)
	GetDataAuthSession(sessionId uuid.UUID) (uuid.UUID, error)
	DeleteAuthSession(sessionId uuid.UUID) error
}

type authSession struct {
	sessionStore  *redis.Storage
	sessionExpiry time.Duration
}

func NewAuthSession() AuthSession {
	sessionConfig := config.LoadAuthSessionConfig()

	return &authSession{
		sessionStore: redis.New(redis.Config{
			Host:     sessionConfig.Host,
			Port:     sessionConfig.Port,
			Username: sessionConfig.Username,
			Reset:    sessionConfig.Reset,
			Database: sessionConfig.Database,
		}),
		sessionExpiry: sessionConfig.Expiration,
	}
}

func (s *authSession) VerifyAuthSession(sessionId uuid.UUID) error {
	userAuthSessionRaw, err := s.sessionStore.Get(sessionId.String())
	if err != nil {
		return err
	}

	userAuthSession, err := uuid.FromBytes(userAuthSessionRaw)
	if err != nil {
		return err
	}

	if userAuthSession != sessionId {
		return errors.New("invalid session id")
	}

	return nil
}

func (s *authSession) CreateAuthSession(citizen models.CitizenAuth) (uuid.UUID, time.Duration, error) {
	binaryUserId, err := citizen.ID.MarshalBinary()
	if err != nil {
		return uuid.Nil, 0, err
	}

	newSessionId := uuid.New()

	if err := s.sessionStore.Set(newSessionId.String(), binaryUserId, s.sessionExpiry); err != nil {
		return uuid.Nil, 0, err
	}

	return newSessionId, s.sessionExpiry, nil
}

func (s *authSession) GetDataAuthSession(sessionId uuid.UUID) (uuid.UUID, error) {
	result, err := s.sessionStore.Get(sessionId.String())
	if err != nil {
		return uuid.Nil, err
	}

	sessionUuid, err := uuid.FromBytes(result)
	if err != nil {
		return uuid.Nil, err
	}

	return sessionUuid, nil
}

func (s *authSession) DeleteAuthSession(sessionId uuid.UUID) error {
	return s.sessionStore.Delete(sessionId.String())
}
