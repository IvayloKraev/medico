package auth

import "github.com/redis/go-redis/v9"

type sessionCommands interface {
}

type session struct {
	store *redis.Client
}

func newSession() *session {
	return &session{
		store: redis.NewClient(&redis.Options{}),
	}
}
