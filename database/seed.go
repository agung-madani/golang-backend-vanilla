package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Seed function inserts sample car data
func Seed() {
	_ = godotenv.Load()

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("SSL_MODE"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	seedSQL := `
	INSERT INTO cars (brand, model, year, price) VALUES
		('Toyota', 'Corolla', 2022, 20000.00),
		('Honda', 'Civic', 2023, 25000.00),
		('Ford', 'Mustang', 2021, 40000.00),
		('Tesla', 'Model 3', 2024, 45000.00)
	ON CONFLICT DO NOTHING;`

	_, err = db.Exec(seedSQL)
	if err != nil {
		log.Fatal("Error inserting seed data:", err)
	}

	fmt.Println("Seeding complete: Sample car data inserted successfully.")
}
