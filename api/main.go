package main

import (
	"booknook/api/database" // import the database package
	"booknook/api/routes"   // import the routes package
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux" // import the mux package
)

func main() {
	// load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// initialize the database connection
	database.InitDB()

	// create a new router instance
	router := mux.NewRouter()

	// set up the router with the api routes
	routes.DefineRoutes(router)
	// health check endpoint
	routes.HealthCheckEndpoint(router)
	// start the http server
	StartServer(router)
}

// * start the http server
func StartServer(router *mux.Router) {
	log.Println("\n\033[1;37;1m * Starting the HTTP server on port:	  ➮\033[1;94;1m 8080\033[0m")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("\n * Failed to start HTTP server: %s\n", err)
	}
}
