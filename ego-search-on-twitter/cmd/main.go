package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	godotenv.Load()

	port := "9999"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
