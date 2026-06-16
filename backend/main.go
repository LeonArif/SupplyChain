package main

import (
	"log"
	"os"

	"supplychain/backend/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := api.SetupRouter()

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
