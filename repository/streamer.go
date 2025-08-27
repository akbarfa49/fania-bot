package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Streamer struct {
	pgx *pgxpool.Pool
}

func NewStreamer(pgx *pgxpool.Pool) *Streamer {
	return &Streamer{pgx: pgx}
}
