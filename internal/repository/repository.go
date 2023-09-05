package repository

import (
	"database/sql"
	"jhonidev/go/go-movie-catalog/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
