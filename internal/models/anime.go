package models

import (
	"time"

	"github.com/google/uuid"
)

type Anime struct {
	ID           uuid.UUID //example 53aa35c8-e659-44b2-882f-f6056e443c99
	Created_at   time.Time
	Updated_at   time.Time
	Deleted_at   time.Time
	Enable       bool
	Language     string
	Title        string
	Genre        string
	EpisodeCount int
	DateRelease  time.Time
	Subtitles    []Subtitles
}

type Subtitles struct {
	ID         uuid.UUID //example 53aa35c8-e659-44b2-882f-f6056e443c99
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
	Enable     bool
	Language   string
	Author     string
}
