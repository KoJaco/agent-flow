// Local dev server: same handler, plain net/http.
package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"example.com/app/internal/app"
)

func main() {
	h, err := app.NewHandler(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	addr := ":" + envOr("PORT", "8080")
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, h))
}

func envOr(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
