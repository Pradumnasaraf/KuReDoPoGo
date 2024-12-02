package main

import (
	"database/sql"
	"log"
	"os"

	random "github.com/Pallinder/go-randomdata"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Using system environment variables")
	} else {
		log.Println("Using .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}

	// Create a product table
	createProductTable(db)

	// Insert a product
	product := Product{
		Name:      random.SillyName(),
		Price:     random.Decimal(10, 100),
		Available: random.Boolean(),
	}
	pk := insertProduct(db, product)
	log.Println("Product created with id:", pk)

	// Get a Product

	//getProducts(db)

	// Get all Products

	//getAllProducts(db)
}

func createProductTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS product (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6, 2) NOT NULL,
		available BOOLEAN,
		created Timestamp DEFAULT NOW()
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, available) VALUES ($1, $2, $3) RETURNING id`
	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func getProducts(db *sql.DB) {
	var name string
	var price float64
	var available bool

	query := `SELECT name, price, available FROM product WHERE id = $1`
	pk := 1
	err := db.QueryRow(query, pk).Scan(&name, &price, &available)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No product found with id:", pk)
		}
		log.Fatal(err)
	}
	log.Println("Product:", name, price, available)
}

func getAllProducts(db *sql.DB) {
	date := []Product{}
	rows, err := db.Query("SELECT name, price, available FROM product")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var price float64
		var available bool
		err := rows.Scan(&name, &price, &available)
		if err != nil {
			log.Fatal(err)
		}
		date = append(date, Product{name, price, available})
	}
	log.Println(date)
}
