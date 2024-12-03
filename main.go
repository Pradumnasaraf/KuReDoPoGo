package main

import (
	"log"
	"os"

	"github.com/Pradumnasaraf/kuredopogo/config"
	"github.com/Pradumnasaraf/kuredopogo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := config.ConnectDB()
	defer db.Close()

	router := gin.Default()
	routes.RegisterRoutes(router, db)
	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
