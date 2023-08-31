package services

import (
	"avito/internal/config"
	"avito/internal/database"
	"avito/internal/models"
	"encoding/json"
	"fmt"
	_ "gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

type GetUser struct {
	UserID string `json:"user_id"`
}

type AddSegmentsToUserStruct struct {
	UserID        string `json:"user_id"`
	AddSegment    string `json:"add_name"`
	DeleteSegment string `json:"delete_name"`
	ExpiresAt     string `json:"expires_at"`
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "Create user"
// @Success 200 {object} models.User
// @Failure 400 {object} error
// @Router /user [post]
func CreateUser(w http.ResponseWriter, _ *http.Request) {
	fmt.Printf("POST /users\n")
	var user models.User
	id, err := database.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create User", http.StatusBadRequest)
		return
	}
	b, err := json.Marshal(id)
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(b)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users without data payload
// @Tags users
// @Produce  json
// @Success 200 {object} []models.User
// @Failure 204 {object} error
// @Router /user/all [get]
func GetUsers(w http.ResponseWriter, _ *http.Request) {
	fmt.Printf("GET /users\n")
	users, errDb := database.GetUsers()
	if errDb != nil {
		http.Error(w, "Users not found", http.StatusNoContent)
	}
	b, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(b)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

// GetUserWithSegments godoc
// @Summary Provides all user segments
// @Description Provides all user segments by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400 {object} error
// @Param user body models.User true "Get user segments"
// @Router /user/segment [post]
func GetUserWithSegments(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET /user\n")

	var user GetUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	segments, err := database.GetUserWithSegments(user.UserID)
	if err != nil {
		log.Println(err)
	}
	var segmentsName []string
	for _, v := range segments {
		if v.Name != "" {
			segmentsName = append(segmentsName, v.Name)
		}
	}
	b, err := json.Marshal(segmentsName)
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(b)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

// AddSegmentsToUser godoc
// @Summary Adds segments to the user
// @Description Provides or deletes all user segments by ID and name of segments
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "Add or delete segments to user"
// @Success 200 {object} string
// @Failure 400 {object} error
// @Router /user/addSegment [post]
func AddSegmentsToUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("POST /users/addSegment\n")

	var req AddSegmentsToUserStruct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	userID := req.UserID
	segmentsAdd := strings.Split(req.AddSegment, ",")
	segmentsDelete := strings.Split(req.DeleteSegment, ",")

	if userID == "" || len(segmentsAdd) == 0 {
		http.Error(w, "Missing user_id or add_name parameter", http.StatusNoContent)
		return
	}
	tx := config.DB.Begin()

	for _, v := range segmentsAdd {
		fmt.Println(v)
		if req.ExpiresAt == "" {
			err = database.AddSegmentToUser(tx, v, userID)
		} else {
			err = database.AddSegmentToUserWithExpiredTime(tx, v, userID, req.ExpiresAt)
		}

		if err != nil {
			tx.Rollback()
			http.Error(w, "Association error", http.StatusBadRequest)
			return
		}
	}
	errTx := tx.Commit().Error
	if errTx != nil {
		fmt.Println(errTx)
		http.Error(w, "Error committing transaction", http.StatusInternalServerError)
		return
	}
	if req.DeleteSegment != "" {
		tx = config.DB.Begin()
		for _, v := range segmentsDelete {
			err = database.DeleteSegmentFromUser(tx, v, userID)
			if err != nil {
				tx.Rollback()
				http.Error(w, "User doesn't have this segment", http.StatusBadRequest)
				return
			}
		}
		errTx = tx.Commit().Error
		if errTx != nil {
			fmt.Println(errTx)
			http.Error(w, "Error committing transaction", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Success"))
	if err != nil {
		log.Println(err)
	}
}
