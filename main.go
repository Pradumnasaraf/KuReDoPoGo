package main

import (
	"log"
	"os"

	"github.com/Pradumnasaraf/kuredopogo/config"
	"github.com/Pradumnasaraf/kuredopogo/middleware"
	"github.com/Pradumnasaraf/kuredopogo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := config.ConnectPostgres()
	defer db.Close()

	middleware.RedisInit()
	defer middleware.RedisClose()

	router := gin.Default()
	router.Use(middleware.RedisRateLimiter())
	routes.RegisterRoutes(router, db)
	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
