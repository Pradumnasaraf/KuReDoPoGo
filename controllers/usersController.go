package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Pradumnasaraf/kuredopogo/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rows, err := db.Query(`SELECT * FROM users`)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while fetching users"})
		}
		defer rows.Close()

		users := []models.User{}
		for rows.Next() {
			var u models.User
			err := rows.Scan(&u.ID, &u.Name, &u.Email)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while fetching users"})
			}
			users = append(users, u)
		}

		ctx.JSON(http.StatusOK, users)
	}
}

func GetUserById(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var u models.User
		id := ctx.Param("id")
		err := db.QueryRow(`SELECT * FROM users WHERE id = $1`, id).Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while fetching user"})
		}
		ctx.JSON(http.StatusOK, u)
	}
}

func CreateUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.ContentType() != "application/json" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Content-Type must be application/json"})
			return
		}
		if ctx.Request.Body == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Request body is empty"})
			return
		}

		var u models.User
		err := ctx.BindJSON(&u)
		if err != nil || u.Name == "" || u.Email == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
			return
		}
		err = db.QueryRow(`INSERT INTO users (name, email) VALUES($1, $2) RETURNING id`, u.Name, u.Email).Scan(&u.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while creating user"})
		}
		ctx.JSON(http.StatusCreated, u)
	}
}

func UpdateUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.ContentType() != "application/json" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Content-Type must be application/json"})
			return
		}
		if ctx.Request.Body == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Request body is empty"})
			return
		}
		var u models.User
		intID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}
		u.ID = intID

		err = ctx.BindJSON(&u)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
			return
		}

		_, execErr := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, u.ID)
		if execErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while updating user"})
		}
		ctx.JSON(http.StatusCreated, u)
	}
}

func DeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while deleting user"})

		}
		ctx.JSON(http.StatusCreated, fmt.Sprintf("User deleted with the ID: %s", id))
	}
}
