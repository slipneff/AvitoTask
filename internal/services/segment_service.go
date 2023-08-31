package services

import (
	"avito/internal/database"
	"avito/internal/models"
	"encoding/json"
	"fmt"
	_ "gorm.io/gorm"
	"net/http"
)

// CreateSegment godoc
// @Summary Create a new segment
// @Description Create a new segment with the input payload
// @Tags Segments
// @Accept  json
// @Produce  json
// @Param Segment body models.Segment true "Create segment"
// @Success 200 {object} models.Segment
// @Failure 400 {object} error
// @Router /segment [post]
func CreateSegment(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("POST /segments\n")
	var segment models.Segment
	err := json.NewDecoder(r.Body).Decode(&segment)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	id, err := database.CreateSegment(segment)
	if err != nil {
		http.Error(w, "Invalid segment", http.StatusBadRequest)
		return
	}
	if segment.Percentage > 0 {
		err = database.RandomApply(
			segment.Percentage,
			segment.Name,
		)
		if err != nil {
			fmt.Println(err)
			http.Error(w, error.Error(err), http.StatusBadRequest)
			return
		}
	}
	b, err := json.Marshal(id)
	w.Write(b)
}

// DeleteSegment godoc
// @Summary Delete a segment
// @Description Delete a segment with the input payload
// @Tags Segments
// @Accept  json
// @Produce  json
// @Param Segment body models.Segment true "Delete segment"
// @Success 200 {object} models.Segment
// @Failure 400 {object} error
// @Router /segment [delete]
func DeleteSegment(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DELETE /segments\n")
	var segment models.Segment
	err := json.NewDecoder(r.Body).Decode(&segment)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	if err := database.DeleteSegment(segment); err != nil {
		http.Error(w, "Segment data not deleted", http.StatusBadRequest)
		return
	}
	w.Write([]byte("Success"))
}
