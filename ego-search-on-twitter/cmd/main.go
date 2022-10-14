package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	f "github.com/kmtym1998/ego-search-twitter"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	godotenv.Load()

	ctx := context.Background()
	if err := funcframework.RegisterEventFunctionContext(ctx, "/", f.Run); err != nil {
		log.Fatalf("funcframework.RegisterEventFunctionContext: %v\n", err)
	}

	port := "9999"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
