package main

import (
	"context"
	"log"

	"github.com/ugabiga/swan/bootstrap/internal/app"
)

// @title		BOOTSTRAP_PLACEHOLDER
// @host		localhost:8080
// @version	0.1
func main() {
	ctx := context.Background()

	newApp := app.NewApp()
	if err := newApp.Start(ctx); err != nil {
		log.Fatal(err)
	}
	if err := newApp.Stop(ctx); err != nil {
		log.Fatal(err)
	}
}
