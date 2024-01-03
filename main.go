// main.go
package main

import (
	"backend-edu-experience/route"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// load latest env

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		return
	}

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	address := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPass, dbHost, dbPort, dbname)

	db, err := gorm.Open(mysql.Open(address), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	router := gin.Default()

	route.InitializeRoutes(router, db)

	// Start the server
	port := os.Getenv("PORT")

	router.GET("/ping", func(ctx *gin.Context) {
		type PingResp struct {
			Data    interface{} `json:"data"`
			Message string      `json:"message"`
		}

		resp := PingResp{
			Data:    nil,
			Message: "success",
		}

		ctx.JSON(201, resp)
	})

	router.Run(":" + port)
}
