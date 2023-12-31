// main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// load env
	err := godotenv.Load(".env") // Load environment variables from .env file
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		return
	}

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	address := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPass, dbHost, dbPort, dbname)

	fmt.Printf("%s", address)
	// Connect to the PostgreSQL database
	db, err := sql.Open("mysql", address)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	// Routes
	// router.GET("/todos", getTodos(db))
	// router.GET("/todos/:id", getTodoByID(db))
	// router.POST("/todos", createTodo(db))
	// router.PUT("/todos/:id", updateTodo(db))
	// router.DELETE("/todos/:id", deleteTodo(db))

	// Start the server
	port := os.Getenv("PORT")

	fmt.Println("halo maria")
	router.Run(":" + port)
}

// func getTodos(db *sql.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		rows, err := db.Query("SELECT id, title, completed FROM todo")
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		defer rows.Close()

// 		todos := []Todo{}
// 		for rows.Next() {
// 			var todo Todo
// 			if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 				return
// 			}
// 			todos = append(todos, todo)
// 		}

// 		c.JSON(http.StatusOK, todos)
// 	}
// }

// Define other CRUD handlers (getTodoByID, createTodo, updateTodo, deleteTodo) similarly
