// main.go
package main

import (
	"backend-edu-experience/route"
	"backend-edu-experience/template"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// load latest env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		return
	}

	sentryDsn := os.Getenv("SENTRY_DSN")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUsername, dbPass, dbname, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	router := gin.Default()

	// this part for sentry

	err = sentry.Init(sentry.ClientOptions{
		Dsn:   sentryDsn,
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
		return
	}

	defer sentry.Flush(2 * time.Second)

	// intialize routes
	route.InitializeRoutes(router, db)

	// Start the server
	port := os.Getenv("PORT")

	router.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(201, template.Response{
			Data:    nil,
			Message: "PONG!!!",
			Error:   nil,
		})
	})

	router.Run(":" + port)
}
