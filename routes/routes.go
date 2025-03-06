package routes

import (
	"database/sql"

	"go_tutor/controllers"

	"github.com/gorilla/mux"
)

// SetupRoutes initializes all routes
func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/cars", controllers.GetCarsHandler(db)).Methods("GET")

	return r
}
