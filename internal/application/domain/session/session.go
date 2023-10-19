package dmsession

import "time"

type Session struct {
	ID                   string    `json:"ID"`
	UserID               int32     `json:"user_id"`
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}
