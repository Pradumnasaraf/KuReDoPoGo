package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Loading the Environment Variables from the system")
	} else {
		log.Print("Loading the Environment Variables from the .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()
	router.GET("/users", getUsers(db))
	router.GET("/users/:id", getUserById(db))
	router.POST("/users", createUser(db))
	router.PUT("/users/:id", updateUser(db))
	router.DELETE("/users/:id", deleteUser(db))

	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}

func getUsers(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rows, err := db.Query(`SELECT * FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		users := []User{}
		for rows.Next() {
			var u User
			err := rows.Scan(&u.ID, &u.Name, &u.Email)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}

		ctx.JSON(http.StatusOK, users)
	}
}

func getUserById(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var u User
		id := ctx.Param("id")
		fmt.Println(id)
		err := db.QueryRow(`SELECT * FROM users WHERE id = $1`, id).Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
				return
			}
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, u)
	}
}

func createUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func updateUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func deleteUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}
