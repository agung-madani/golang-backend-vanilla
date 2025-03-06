package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"go_tutor/models"
)

// GetCarsHandler handles the GET /cars request
func GetCarsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cars, err := models.GetAllCars(db)
		if err != nil {
			http.Error(w, "Failed to fetch cars", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cars)
	}
}
