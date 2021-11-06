package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	Host string
	DbPort int
	User string
	Password string
	Dbname string
	ServerPort string
	StripeKey string
}

func GetEnvVars() *Env {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("app: Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	serverPort := os.Getenv("SERVER_PORT")
	stripeKey := os.Getenv("STRIPE_SECRET_KEY")

	parsedPort, err := strconv.Atoi(port)

	if err != nil {
		log.Fatal(err.Error())
	}

	return &Env{
			Host: host,
			DbPort: parsedPort,
			User: user, 
			Password: password,
			Dbname: dbname,
			ServerPort: serverPort,
			StripeKey: stripeKey,
	}
}
