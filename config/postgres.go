package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Loading the Environment Variables from the system")
	} else {
		log.Print("Loading the Environment Variables from the .env file")
	}
}

func ConnectPostgres() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URI"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id BIGSERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(50) NOT NULL
	);`)
	if err != nil {
		log.Fatal("Failed to ensure table exists:", err)
	}

	return db
}
