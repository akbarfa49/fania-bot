package repository

import (
	"context"
	"errors"
	"fania-bot/model"
	"time"

	"github.com/jackc/pgx/v5"
)

func (s *Streamer) FindNotificationByUserIdentifier(ctx context.Context, userIdentifier, streamPlatform string) ([]model.StreamNotificationTable, error) {

	out := make([]model.StreamNotificationTable, 0, 20)
	const sql = `SELECT platform, guild, channel, message, created_at, updated_at, metadata
		FROM stream_notification
		WHERE user_identifier = $1
		AND stream_platform = $2`
	rows, err := s.pgx.Query(ctx, sql, userIdentifier, streamPlatform)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return out, nil
		}
		return out, err
	}
	for rows.Next() {
		tmp := model.StreamNotificationTable{UserIdentifier: userIdentifier, StreamPlatform: streamPlatform}
		if err := rows.Scan(&tmp.Platform, &tmp.Guild, &tmp.Channel, &tmp.Message, &tmp.CreatedAt, &tmp.UpdatedAt, &tmp.Metadata); err != nil {
			return nil, err
		}
		out = append(out, tmp)
	}

	return out, nil

}

func (s *Streamer) FindActiveNotificationByUserIdentifierAndStreamPlatform(ctx context.Context, userIdentifier, streamPlatform string) ([]model.StreamNotificationTable, error) {

	out := make([]model.StreamNotificationTable, 0, 20)
	const sql = `SELECT platform, guild, channel, message, created_at, updated_at, metadata
		FROM stream_notification
		WHERE user_identifier = $1
		AND stream_platform = $2
		AND deleted_at is null`
	rows, err := s.pgx.Query(ctx, sql, userIdentifier, streamPlatform)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return out, nil
		}
		return out, err
	}
	for rows.Next() {
		tmp := model.StreamNotificationTable{UserIdentifier: userIdentifier, StreamPlatform: streamPlatform}
		if err := rows.Scan(&tmp.Platform, &tmp.Guild, &tmp.Channel, &tmp.Message, &tmp.CreatedAt, &tmp.UpdatedAt, &tmp.Metadata); err != nil {
			return nil, err
		}
		out = append(out, tmp)
	}

	return out, nil

}

func (s *Streamer) FindActiveNotificationStreamPlatformAndUserIdentifierGroupByUserIdentifier(ctx context.Context) ([]model.StreamNotificationTable, error) {

	out := make([]model.StreamNotificationTable, 0, 20)
	const sql = `SELECT user_identifier, stream_platform
		FROM stream_notification
		WHERE  deleted_at is null
		GROUP BY user_identifier,stream_platform`
	rows, err := s.pgx.Query(ctx, sql)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return out, nil
		}
		return out, err
	}
	for rows.Next() {
		tmp := model.StreamNotificationTable{}
		if err := rows.Scan(&tmp.UserIdentifier, &tmp.StreamPlatform); err != nil {
			return nil, err
		}
		out = append(out, tmp)
	}

	return out, nil

}

func (s *Streamer) FindActiveNotificationStreamPlatformAndUserUniqueIDGroupByUserIdentifier(ctx context.Context) ([]model.StreamNotificationTable, error) {

	out := make([]model.StreamNotificationTable, 0, 20)
	const sql = `SELECT user_identifier, user_unique_id, stream_platform
		FROM stream_notification
		WHERE  deleted_at is null
		GROUP BY user_unique_id,stream_platform,user_identifier`
	rows, err := s.pgx.Query(ctx, sql)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return out, nil
		}
		return out, err
	}
	for rows.Next() {
		tmp := model.StreamNotificationTable{}
		if err := rows.Scan(&tmp.UserIdentifier, &tmp.UserUniqueID, &tmp.StreamPlatform); err != nil {
			return nil, err
		}
		out = append(out, tmp)
	}

	return out, nil

}

func (s *Streamer) CreateNotification(ctx context.Context, data model.StreamNotificationTable) error {

	const sql = `INSERT INTO stream_notification (platform, guild, channel, user_identifier, user_unique_id, stream_platform, message, created_at, updated_at, metadata)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
	`
	_, err := s.pgx.Exec(ctx, sql, data.Platform, data.Guild, data.Channel, data.UserIdentifier, data.UserUniqueID, data.StreamPlatform, data.Message, data.CreatedAt, data.UpdatedAt, data.Metadata)
	return err

}

func (s *Streamer) CreateNotifDelay(ctx context.Context, key string, expiredAt time.Time) error {
	const sql = ` INSERT INTO stream_notif_delay (key, expired_at) VALUES ($1, $2)`
	_, err := s.pgx.Exec(ctx, sql, key, expiredAt)
	return err
}

func (s *Streamer) NotifDelayIsActive(ctx context.Context, key string) bool {
	const sql = `select 1 from stream_notif_delay where key = $1;`
	var i int
	err := s.pgx.QueryRow(ctx, sql, key).Scan(&i)
	if err == nil {
		return true
	}
	return !errors.Is(err, pgx.ErrNoRows)
}
