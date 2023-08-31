package services_test

import (
	"avito/internal/config"
	"avito/internal/models"
	db "avito/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	_ "time"

	"avito/internal/services"
)

func TestGenerateReportCSV(t *testing.T) {
	r := setupRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Создаем канал для синхронизации
	done := make(chan struct{})

	go func() {
		defer close(done)

		reqBody := services.HistoryStruct{
			UserID: "testUserID",
			Year:   2023,
			Month:  8,
		}
		reqJSON, err := json.Marshal(reqBody)
		assert.NoError(t, err)

		res, err := http.Post(ts.URL+"/history", "application/json", bytes.NewBuffer(reqJSON))
		assert.NoError(t, err)
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)

		assert.Equal(t, "attachment; filename=report.csv", res.Header.Get("Content-Disposition"))
		assert.Equal(t, "text/csv", res.Header.Get("Content-Type"))
	}()

	// Ждем завершения обработки запроса
	<-done
}

// Остальной код остается без изменений

func TestCreateSegment(t *testing.T) {
	r := setupRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	segment := models.Segment{
		Name:       "TestSegment",
		Percentage: 10,
	}
	reqJSON, err := json.Marshal(segment)
	assert.NoError(t, err)

	res, err := http.Post(ts.URL+"/segment", "application/json", bytes.NewBuffer(reqJSON))
	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestDeleteSegment(t *testing.T) {

	r := setupRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	segment := models.Segment{
		Name: "TestSegment",
	}
	reqJSON, err := json.Marshal(segment)
	assert.NoError(t, err)

	req, err := http.NewRequest("DELETE", ts.URL+"/segment", bytes.NewBuffer(reqJSON))
	assert.NoError(t, err)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCreateUser(t *testing.T) {
	r := setupRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	user := models.User{}
	reqJSON, err := json.Marshal(user)
	assert.NoError(t, err)

	res, err := http.Post(ts.URL+"/user", "application/json", bytes.NewBuffer(reqJSON))
	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetUsers(t *testing.T) {
	r := setupRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/user/all")
	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetUserWithSegments(t *testing.T) {
	r := setupRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	user := services.GetUser{
		UserID: "1",
	}
	reqJSON, err := json.Marshal(user)
	assert.NoError(t, err)

	res, err := http.Post(ts.URL+"/user", "application/json", bytes.NewBuffer(reqJSON))
	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestAddSegmentsToUser(t *testing.T) {
	r := setupRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	addSegments := services.AddSegmentsToUserStruct{
		UserID:     "1",
		AddSegment: "Segment1,Segment2",
	}
	reqJSON, err := json.Marshal(addSegments)
	assert.NoError(t, err)

	res, err := http.Post(ts.URL+"/user/addSegment", "application/json", bytes.NewBuffer(reqJSON))
	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func setupRouter() http.Handler {
	err := godotenv.Load(".env")
	config.Setup()
	fmt.Println("DB Connected")
	err = config.DB.AutoMigrate(&db.User{}, &db.Segment{}, &db.SegmentHistory{}, &db.UserSegments{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	r := mux.NewRouter()
	r.HandleFunc("/user", services.CreateUser).Methods("POST")
	r.HandleFunc("/user/all", services.GetUsers).Methods("GET")
	r.HandleFunc("/user", services.GetUserWithSegments).Methods("GET")
	r.HandleFunc("/user/addSegment", services.AddSegmentsToUser).Methods("POST")
	r.HandleFunc("/history", services.GenerateReportCSV).Methods("POST")
	r.HandleFunc("/segment", services.CreateSegment).Methods("POST")
	r.HandleFunc("/segment", services.DeleteSegment).Methods("DELETE")
	return r
}
