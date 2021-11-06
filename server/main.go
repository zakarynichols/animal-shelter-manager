package main

import (
	"fmt"
	"log"
	"net/http"
	"server/app"
	"server/config"
	"server/database"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	stripe "github.com/stripe/stripe-go/v72"
)

  func main() {
	// Get environment variables
	env := config.GetEnvVars()

	// Initialize database
	store, err := database.OpenDB(database.DatabaseInfo{
		Host: env.Host,
		Port: env.DbPort,
		User: env.User,
		Password: env.Password,
		Dbname: env.Dbname,
	})
	
	// Database is a required dependency. Terminate on error
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create new router instance
	router := mux.NewRouter()

	// Instantiate new server with required dependencies
	server := app.NewServer(store, router, env)

	// Initialize stripe API
	stripe.Key = env.StripeKey

	// Initialize routes
	server.InitializeRoutes()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", env.ServerPort), server.Router()))
  }