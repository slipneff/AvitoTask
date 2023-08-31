package models

type Segment struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	Percentage int
	Users      []*User `gorm:"many2many:user_segments;"`
}
