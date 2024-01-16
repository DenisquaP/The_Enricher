package postgres

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func NewConfig() Config {
	err := godotenv.Load(".env", ".env.example")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	if host == "" {
		log.Fatal("can`t find a Database host")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("want int, got err: %v", err)
	}

	user := os.Getenv("USER")
	if user == "" {
		log.Fatal("can`t find a Database user")
	}

	password := os.Getenv("PASSWORD")
	if password == "" {
		log.Fatal("can`t find a Database password")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("can`t find a Database name")
	}
	return Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DbName:   dbName,
	}
}
