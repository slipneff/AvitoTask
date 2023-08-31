package tools

import (
	"avito/internal/models"
	"reflect"
)

const (
	AddSegment    string = "Add"
	DeleteSegment        = "Delete"
)

func Contains(s []*models.Segment, e models.Segment) bool {
	for _, v := range s {
		if reflect.DeepEqual(v, e) {
			return true
		}
	}
	return false
}
