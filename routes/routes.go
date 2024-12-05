package routes

import (
	"database/sql"

	"github.com/Pradumnasaraf/kuredopogo/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	
	router.GET("/health", controllers.HealthCheck())

	router.GET("/users", controllers.GetUsers(db))
	router.GET("/users/:id", controllers.GetUserById(db))
	router.POST("/users", controllers.CreateUser(db))
	router.PUT("/users/:id", controllers.UpdateUser(db))
	router.DELETE("/users/:id", controllers.DeleteUser(db))
}
