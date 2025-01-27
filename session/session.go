package session

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"medico/models"
	"time"
)

type AuthSession interface {
	VerifyAuthSession(sessionId uuid.UUID) error
	CreateAuthSession(citizen models.Citizen) (error, string)
	GetDataAuthSession(sessionId uuid.UUID) authSessionData // TODO change to models.Citizen or something like that
	DeleteAuthSession(sessionId uuid.UUID) error
}

type authSessionData struct {
	userId    models.ModelID
	createdAt time.Time
}

type authSession struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewSession() AuthSession {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "default",
	})

	ctx := context.Background()

	return &authSession{
		redisClient: client,
		ctx:         ctx,
	}
}

func (s *authSession) VerifyAuthSession(sessionId uuid.UUID) error {
	userAuthSession, err := s.redisClient.Get(s.ctx, sessionId.String()).Result()

	if err != nil {
		return err
	}

	if userAuthSession == "" {
		return errors.New("invalid session")
	}

	return nil
}

func (s *authSession) CreateAuthSession(citizen models.Citizen) (error, string) {
	newAuthSessionData := &authSessionData{
		userId:    citizen.ID,
		createdAt: time.Now(),
	}

	authSessionDataMarshaled, err := json.Marshal(newAuthSessionData)
	if err != nil {
		return err, ""
	}

	newSessionId := uuid.New().String()

	return s.redisClient.Set(s.ctx, newSessionId, authSessionDataMarshaled, 1000).Err(), newSessionId
}

func (s *authSession) GetDataAuthSession(sessionId uuid.UUID) authSessionData {
	result, err := s.redisClient.Get(s.ctx, sessionId.String()).Result()
	if err != nil {
		return authSessionData{}
	}

	var receivedAuthSessionData authSessionData
	err = json.Unmarshal([]byte(result), &receivedAuthSessionData)

	if err != nil {
		return authSessionData{}
	}

	return receivedAuthSessionData
}

func (s *authSession) DeleteAuthSession(sessionId uuid.UUID) error {
	return s.redisClient.Del(s.ctx, sessionId.String()).Err()
}
