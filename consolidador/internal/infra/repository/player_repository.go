package repository

import (
	"consolidador/internal/infra/db"
	"database/sql"
)

type PlayerRepository struct {
	Repository
}

func NewPlayerRepository(dbConn *sql.DB) *PlayerRepository {
	return &PlayerRepository{
		Repository: Repository{
			dbConn:  dbConn,
			Queries: db.New(dbConn),
		},
	}
}
