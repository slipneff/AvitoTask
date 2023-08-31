package models

import (
	"gorm.io/gorm"
	"time"
)

type SegmentHistory struct {
	gorm.Model
	UserID      string
	SegmentName string
	Operation   string
	Timestamp   time.Time
}
