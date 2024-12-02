package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Using system environment variables")
	} else {
		log.Println("Using .env file")
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey, I'm a server!",
		})
	})
	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
