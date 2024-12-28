package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		log.Fatal("PORT  was not found in enviroment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1router := chi.NewRouter()
	v1router.Get("/healthy", handlerReadiness)
	v1router.Get("/err", handleError)

	router.Mount("/v1", v1router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portNumber,
	}

	log.Printf("Server started on port %v", portNumber)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server had error")
	}

	fmt.Printf("Port number: %s", portNumber)
}
