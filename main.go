package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	DotEnvError := godotenv.Load(".env")
	if DotEnvError != nil {
		log.Fatalf("Error loading .env file.")
		return
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT must be set in the environment")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	log.Printf(`Server starting on port %s`, portString)
	ServeError := srv.ListenAndServe()
	if ServeError != nil {
		log.Fatalf(`Error starting server : %s`, ServeError)
	}

	log.Printf("Listening on port %s", portString)
}
