package user

import "time"

type Service struct {
}

type UserAuth struct {
	Source           string    `json:"source"`
	SourceIdentifier string    `json:"source_identifier"`
	UserID           string    `json:"user_id"`
	LinkedAt         time.Time `json:"linked_at"`
}

type User struct {
	ID              string    `json:"id"`
	Username        string    `json:"username"`
	ProfileImageURL string    `json:"profile_image_url"`
	CreatedAt       time.Time `json:"created_at"`
}
