package repository

import (
	"context"
	"errors"
	"fania-bot/model"

	"github.com/jackc/pgx/v5"
)

/*
CREATE TABLE stream_history (

	id SERIAL PRIMARY KEY,
	platform VARCHAR(255),
	user_identifier VARCHAR(255),
	stream_identifier VARCHAR(255),
	created_at TIMESTAMP

);
*/

func (s *Streamer) FindLatestHistoryByID(ctx context.Context, platform, userIdentifier string) (model.StreamHistoryTable, bool, error) {
	out := model.StreamHistoryTable{}
	const sql = `SELECT platform, user_identifier, stream_identifier, created_at
	FROM stream_history
	WHERE platform = $1
	  AND user_identifier = $2
	  ORDER BY created_at DESC
		LIMIT 1;`
	err := s.pgx.QueryRow(ctx, sql, platform, userIdentifier).Scan(&out.Platform, &out.UserIdentifier, &out.StreamIdentifier, &out.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return out, false, nil
		}
		return out, false, err
	}
	return out, true, nil
}

func (s *Streamer) InsertHistory(ctx context.Context, data model.StreamHistoryTable) error {
	const sql = `INSERT INTO stream_history (platform, user_identifier, stream_identifier, created_at)
VALUES ($1, $2, $3, $4);
`
	_, err := s.pgx.Exec(ctx, sql, data.Platform, data.UserIdentifier, data.StreamIdentifier, data.CreatedAt)
	return err
}
