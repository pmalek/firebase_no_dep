package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
)

func main() {
	if _, err := firebase.NewApp(context.Background(), nil); err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
}
