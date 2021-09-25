package models

import "time"

type Track struct {
	UserID    string    `json:"user_id"`
	TrackID   string    `json:"track_id"`
	Features  string    `json:"features"`
	Mood      int       `json:"mood"`
	CreatedAt time.Time `json:"created_at"`
}
