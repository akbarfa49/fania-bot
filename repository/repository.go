package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type storage struct {
	pgx pgxpool.Conn
}
