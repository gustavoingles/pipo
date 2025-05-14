package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var dbpool *pgxpool.Pool

func main() {
	// Loading the environment variables from a .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load the environment variables: %v", err)
	}

	// Establishing a connection pool to the database
	dbpool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	//Function to set up all routes to some handler
	setupRoutes()

	// Starting the web server by putting it to listen some port in an ip adress an then serve the application
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
