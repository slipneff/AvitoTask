package db_test

import (
	"avito/internal/config"
	"avito/internal/database"
	"avito/internal/models"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	_ "time"
)

type v2Suite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
	user models.User
}

func TestCreateUser(t *testing.T) {
	dsn := "host=localhost user=postgres password=admin dbname=avitodb sslmode=disable"
	s := &v2Suite{}
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  dsn,
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.db, err = gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)
	defer db.Close() // Закрытие соединения с базой данных после завершения теста

	s.user = models.User{}

	id, err := database.CreateUser(s.user)
	assert.NoError(t, err)
	assert.NotEqual(t, uint(0), id)
}

func TestGetUsers(t *testing.T) {
	dsn := "host=localhost user=postgres password=admin dbname=avitodb sslmode=disable"
	s := &v2Suite{}
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  dsn,
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.db, err = gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)
	defer db.Close()
	users := [4]models.User{}
	for _, v := range users {
		id, err := database.CreateUser(v)
		assert.NoError(t, err)
		assert.NotEqual(t, uint(0), id)
	}

	resultUsers, err := database.GetUsers()
	assert.NoError(t, err)
	assert.Equal(t, len(users), len(resultUsers))
}

func TestCreateSegment(t *testing.T) {
	dsn := "host=localhost user=postgres password=admin dbname=avitodb sslmode=disable"
	s := &v2Suite{}
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  dsn,
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.db, err = gorm.Open(dialector, &gorm.Config{})
	assert.NoError(t, err)
	defer db.Close()
	segment := models.Segment{
		Name:       "TEST",
		Percentage: 10,
	}

	id, err := database.CreateSegment(segment)
	assert.NoError(t, err)
	assert.NotEqual(t, uint(0), id)
}

func TestRandomApply(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	oldDB := config.DB
	config.DB = mockDB
	defer func() { config.DB = oldDB }()

	percentage := 10
	segmentName := "test_segment"
	mock.ExpectExec("INSERT INTO \"user_segments\"").WillReturnResult(sqlmock.NewResult(0, 1))

	err = database.RandomApply(percentage, segmentName)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
