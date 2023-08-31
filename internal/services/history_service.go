package services

import (
	"avito/internal/database"
	"encoding/json"
	_ "gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type HistoryStruct struct {
	UserID string `json:"user_id"`
	Year   int    `json:"year"`
	Month  int    `json:"month"`
}

// GenerateReportCSV godoc
// @Summary Generate a history
// @Description Generates a history of requests for a specified time
// @Tags History
// @Accept  json
// @Produce  json
// @Param History body HistoryStruct true "Generate history"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /history [post]
func GenerateReportCSV(w http.ResponseWriter, r *http.Request) {

	var req HistoryStruct
	err := json.NewDecoder(r.Body).Decode(&req)
	reportContent, err := database.GenerateCSVReport(req.UserID, req.Year, time.Month(req.Month))
	if err != nil {
		http.Error(w, "Error generating CSV report", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=report.csv")
	w.Header().Set("Content-Type", "text/csv")
	_, err = w.Write([]byte(reportContent))
	if err != nil {
		log.Println(err)
	}
}
