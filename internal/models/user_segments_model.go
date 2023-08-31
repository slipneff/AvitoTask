package models

import "time"

type UserSegments struct {
	UserID    int
	SegmentID int
	ExpiresAt time.Time
}
