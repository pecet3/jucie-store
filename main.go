package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/pecet3/my-api/auth"
	"github.com/pecet3/my-api/controllers"
	"github.com/pecet3/my-api/data"
	"github.com/pecet3/my-api/storage"
)

type App struct {
	sessions  *auth.SessionStore
	validator *validator.Validate
	data      data.Data
	storage   storage.StorageServices
}

func main() {
	log.Println("Running the server...")
	loadEnv()

	mux := http.NewServeMux()
	data := data.New()
	app := App{
		sessions:  auth.NewSessionStore(data),
		validator: validator.New(),
		data:      data,
		storage:   storage.Services{},
	}

	auth.Run(mux, app.sessions, app.data, app.validator)
	storage.Run(mux, app.data.Db, app.sessions)

	controllers.Run(mux, app.data, app.storage)

	address := "127.0.0.1:8090"
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}
	log.Printf("Server is listening on: [%s]", address)
	log.Fatal(server.ListenAndServe())
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println("Loaded .env")
}
