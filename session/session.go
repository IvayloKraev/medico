package session

import (
	"fmt"
	"github.com/gofiber/storage/redis/v3"
	"github.com/google/uuid"
	"medico/config"
	"time"
)

type AuthSession interface {
	CreateAuthSession(userId uuid.UUID) (uuid.UUID, time.Duration, error)
	GetAuthSession(sessionId uuid.UUID) (uuid.UUID, error)
	DeleteAuthSession(sessionId uuid.UUID) error

	CreateAuthSessionWithSubrole(subrole string, userId uuid.UUID) (uuid.UUID, time.Duration, error)
	GetAuthSessionWithSubrole(subrole string, sessionId uuid.UUID) (uuid.UUID, error)
	DeleteAuthSessionWithSubrole(sessionId uuid.UUID) error
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

func (s *authSession) CreateAuthSession(userId uuid.UUID) (uuid.UUID, time.Duration, error) {
	userIdBytes, err := userId.MarshalBinary()
	if err != nil {
		return uuid.Nil, 0, err
	}

	newSessionId := uuid.New()

	if err := s.sessionStore.Set(fmt.Sprintf("%s:%s", s.role, newSessionId.String()), userIdBytes, s.sessionExpiry); err != nil {
		return uuid.Nil, 0, err
	}

	return newSessionId, s.sessionExpiry, nil
}

func (s *authSession) GetAuthSession(sessionId uuid.UUID) (uuid.UUID, error) {
	userIdBytes, err := s.sessionStore.Get(fmt.Sprintf("%s:%s", s.role, sessionId.String()))
	if err != nil {
		return uuid.Nil, err
	}

	userId, err := uuid.FromBytes(userIdBytes)
	if err != nil {
		return uuid.Nil, err
	}

	return userId, nil
}

func (s *authSession) DeleteAuthSession(sessionId uuid.UUID) error {
	return s.sessionStore.Delete(fmt.Sprintf("%s:%s", s.role, sessionId.String()))
}

func (s *authSession) CreateAuthSessionWithSubrole(subrole string, userId uuid.UUID) (uuid.UUID, time.Duration, error) {
	userIdBytes, err := userId.MarshalBinary()
	if err != nil {
		return uuid.Nil, 0, err
	}

	newSessionId := uuid.New()

	if err := s.sessionStore.Set(fmt.Sprintf("%s:%s:%s", s.role, subrole, newSessionId.String()), userIdBytes, s.sessionExpiry); err != nil {
		return uuid.Nil, 0, err
	}

	return newSessionId, s.sessionExpiry, nil
}

func (s *authSession) GetAuthSessionWithSubrole(subrole string, sessionId uuid.UUID) (uuid.UUID, error) {
	userIdBytes, err := s.sessionStore.Get(fmt.Sprintf("%s:%s:%s", s.role, subrole, sessionId.String()))
	if err != nil {
		return uuid.Nil, err
	}

	userId, err := uuid.FromBytes(userIdBytes)
	if err != nil {
		return uuid.Nil, err
	}

	return userId, nil
}

func (s *authSession) DeleteAuthSessionWithSubrole(sessionId uuid.UUID) error {
	return s.sessionStore.Delete(fmt.Sprintf("%s:*:%s", s.role, sessionId.String()))
}
