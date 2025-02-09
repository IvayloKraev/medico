package session

import (
	"errors"
	"fmt"
	"github.com/gofiber/storage/redis/v3"
	"github.com/google/uuid"
	"medico/config"
	"time"
)

type AuthSession interface {
	VerifyAuthSession(sessionId uuid.UUID) error
	CreateAuthSession(userId uuid.UUID) (uuid.UUID, time.Duration, error)
	GetDataAuthSession(sessionId uuid.UUID) (uuid.UUID, error)
	DeleteAuthSession(sessionId uuid.UUID) error
}

type authSession struct {
	sessionStore  *redis.Storage
	sessionExpiry time.Duration
	role          string
}

func NewAuthSession(role string) AuthSession {
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
		role:          role,
	}
}

func (s *authSession) VerifyAuthSession(sessionId uuid.UUID) error {
	userAuthSession, err := s.GetDataAuthSession(sessionId)
	if err != nil {
		return err
	}

	if userAuthSession != sessionId {
		return errors.New("invalid session id")
	}

	return nil
}

func (s *authSession) CreateAuthSession(userId uuid.UUID) (uuid.UUID, time.Duration, error) {
	binaryUserId, err := userId.MarshalBinary()
	if err != nil {
		return uuid.Nil, 0, err
	}

	newSessionId := uuid.New()

	if err := s.sessionStore.Set(fmt.Sprintf("%s:%s", s.role, newSessionId.String()), binaryUserId, s.sessionExpiry); err != nil {
		return uuid.Nil, 0, err
	}

	return newSessionId, s.sessionExpiry, nil
}

func (s *authSession) GetDataAuthSession(sessionId uuid.UUID) (uuid.UUID, error) {
	result, err := s.sessionStore.Get(fmt.Sprintf("%s:%s", s.role, sessionId.String()))
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
	return s.sessionStore.Delete(fmt.Sprintf("%s:%s", s.role, sessionId.String()))
}
