package main

import (
	"avito/internal/config"
	"avito/internal/database"
	db "avito/internal/models"
	"avito/internal/server"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

// @title           Avito Test Task
// @version         1.0
// @description     This is the swagger document for the Avito task.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Slipneff
// @contact.url    https://github.com/slipneff

// @host      localhost:80
// @BasePath  /

func main() {
	err := godotenv.Load(".env")
	config.Setup()
	fmt.Println("DB Connected")
	err = config.DB.AutoMigrate(&db.User{}, &db.Segment{}, &db.SegmentHistory{}, &db.UserSegments{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("DB Migrated")
	c := cron.New()
	err = c.AddFunc("* /1 * * * *", database.DeleteExpiredSegments)
	c.Start()
	server.NewHTTPServer()
	fmt.Println("Server started")
}
