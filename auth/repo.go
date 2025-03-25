package auth

import (
	"context"
	"github.com/jackc/pgx/v5"
	baseRepo "medico/repo"
	"medico/repo/db"
	"medico/utils"
)

type repoQueries struct {
}

type repoMutations struct {
}

type repo struct {
	connection *pgx.Conn
	queries    *db.Queries
}

func newRepo(ctxOptional ...context.Context) *repo {
	ctx := utils.FirstContextOrBackground(ctxOptional)
	connection := baseRepo.NewConnection(utils.GetDatabaseConfig(), ctx)
	queries := db.New()
	return &repo{
		connection: connection,
		queries:    queries,
	}
}
