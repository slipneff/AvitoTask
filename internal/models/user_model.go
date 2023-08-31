package models

type User struct {
	ID       uint       `gorm:"primaryKey"`
	Segments []*Segment `gorm:"many2many:user_segments;"`
}
