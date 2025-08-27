package model

import (
	"database/sql"
	"time"
)

type StreamHistoryTable struct {
	Platform         string
	UserIdentifier   string
	StreamIdentifier string
	CreatedAt        time.Time
}

type StreamNotificationTable struct {
	StreamPlatform string

	// platform to send
	Platform string

	// each platform have their own format so better use parser
	Guild    string
	Channel  string
	Metadata sql.NullString

	UserIdentifier string
	UserUniqueID   string

	// Message to send
	Message string
	// NextNotif time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type DiscordMetadata struct {
	// Sender       string `json:"sender"`
	WebhookID    string `json:"webhook_id"`
	WebhookToken string `json:"webhook_token"`
	Username     string `json:"username"`
	AvatarURI    string `json:"avatar_uri"`
	// UserToken    string `json:"user_token"`
}

type StreamBroadcasterTable struct {
	StreamPlatform string
	UserIdentifier string
	UserUniqueID   string
	// implement delay for next notif
	NextNotif            time.Time
	LastStreamIdentifier string
	CreatedAt            time.Time
}
