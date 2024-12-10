package db

import (
	"fmt"
	"module/Service/intermal/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	databasePool *pgxpool.Pool
}

func NewStorage(pool *pgxpool.Pool) *Storage {
	storage := New(Storage)
	storage.databasePool = pool
	return storage
}

func (storage *Storage) List(nameFilter string) []models.Models {
	query := "SELECT id, name, value, timestamp FROM models"
	args := make([]interface{}, 0)
	if nameFilter != "" {
		query += "WHERE name LIKE $1"
		args = append(args, fmt.Sprintf("%%%s%%", nameFilter))
	}

	var result []models.Models

	pgxscan.Select(storage.databasePool, &result, query, args)
	return result
}
