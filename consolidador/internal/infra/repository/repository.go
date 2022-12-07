package repository

import (
	"consolidador/internal/infra/db"
	"database/sql"
)

type Repository struct {
	dbConn *sql.DB
	*db.Queries
}
