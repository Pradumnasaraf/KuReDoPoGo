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

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT NOT NULL, email TEXT NOT NULL)`)

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
		var u User
		err := ctx.BindJSON(&u)
		if err != nil {
			log.Fatal(err)
		}
		var id int
		err = db.QueryRow(`INSERT INTO users (name email) VALUES($1, $2) RETURNING id`, u.Name, u.Email).Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusCreated, fmt.Sprintf("User created with the ID: %b", id))
	}
}

func updateUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var u User
		id := ctx.Param("id")
		err := ctx.BindJSON(&u)
		if err != nil {
			log.Fatal(err)
		}

		_, execErr := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, id)
		if execErr != nil {
			log.Fatal(execErr)
		}
		ctx.JSON(http.StatusCreated, fmt.Sprintf("User updated with the ID: %s", id))
	}
}

func deleteUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusCreated, fmt.Sprintf("User updated with the ID: %s", id))
	}
}
