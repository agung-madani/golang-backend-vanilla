package models // Declares the package name

import (
	"database/sql" // Imports the database/sql package for database interactions
	"log"          // Imports the log package for logging errors
	"time"         // Imports the time package for working with time values
)

// Car struct represents the cars table
type Car struct {
    ID        string    `json:"id"`        // Car's unique identifier
    Brand     string    `json:"brand"`     // Car's brand
    Model     string    `json:"model"`     // Car's model
    Year      int       `json:"year"`      // Year the car was manufactured
    Price     float64   `json:"price"`     // Car's price
    CreatedAt time.Time `json:"created_at"` // Timestamp of when the car record was created
}

// GetAllCars retrieves all cars from the database
func GetAllCars(db *sql.DB) ([]Car, error) {
    rows, err := db.Query("SELECT id, brand, model, year, price, created_at FROM cars") // Queries the database to retrieve all cars
    if err != nil {
        return nil, err // Returns an error if the query fails
    }
    defer rows.Close() // Ensures the database rows are closed after the function finishes

    var cars []Car // Initializes a slice to hold the retrieved cars

    for rows.Next() {
        var car Car // Declares a variable to hold each car record
        if err := rows.Scan(&car.ID, &car.Brand, &car.Model, &car.Year, &car.Price, &car.CreatedAt); err != nil {
            log.Println("Error scanning row:", err) // Logs an error if scanning a row fails
            continue // Skips to the next row if an error occurs
        }
        cars = append(cars, car) // Adds the car to the slice
    }

    return cars, nil // Returns the slice of cars and nil as the error
}
