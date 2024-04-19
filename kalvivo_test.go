package klaviyo_test

import (
	"log"
	"os"

	"github.com/bluettipower/klaviyo-go"
	"github.com/joho/godotenv"
)

func setupClient() *klaviyo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	APIkey := os.Getenv("APIKey")
	return klaviyo.NewClient(APIkey)
}
