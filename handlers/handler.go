package handlers

import (
	"database/sql"

	"github.com/clay10j/beerdex/internal/database"
)

type HandlerConfig struct {
	DB        *sql.DB
	DBQueries *database.Queries
}
