package main

import (
	"fmt"
	"log"
	"net/http"
	"server/app"
	"server/config"
	"server/database"
	"server/handlers"
	"server/store"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	stripe "github.com/stripe/stripe-go/v72"
)

func main() {
	// Get environment variables
	env := config.GetEnvVars()

	// Initialize database
	db, err := database.OpenDB(database.DatabaseInfo{
		Host:     env.Host,
		Port:     env.DbPort,
		User:     env.User,
		Password: env.Password,
		Dbname:   env.Dbname,
	})

	// Database is a required dependency. Terminate on error
	if err != nil {
		log.Fatal(err.Error())
	}

	// Instantiate store abstraction for http handlers to interface with db
	store := store.NewStore(db)

	// Instantiate new router handlers
	handlers := handlers.NewHandlers()

	// Instantiate new router instance
	router := mux.NewRouter()

	// Instantiate new server with required dependencies
	server := app.NewServer(store, router, env, handlers)

	// Initialize stripe API
	stripe.Key = env.StripeKey

	// Initialize app service
	server.InitAppService()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", env.ServerPort), router))
}
