package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	JwtSecret  string
	ExpireTime int
)

func init() {

	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Problem loading .env file: %v", err)
		os.Exit(-1)
	}

	JwtSecret = os.Getenv("JWT_SECRET")

	ExpireTime, err = strconv.Atoi(os.Getenv("EXPIRE_TIME"))
	if err != nil {
		log.Fatalf("Problem loading EXPIRE_TIME: %v", err)
		os.Exit(-1)
	}

}
