package main

import (
	"github.com/3110Y/reminder/pkg/Domain/models"
	"github.com/3110Y/reminder/pkg/application/service"
	"github.com/3110Y/reminder/pkg/infrastructure/database"
	"github.com/3110Y/reminder/pkg/infrastructure/repository"
	"github.com/3110Y/reminder/pkg/presentation/http/handler/event"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("APP_PORT")
	router := chi.NewRouter()
	db := database.NewProviderGorm()
	eventRepository := repository.NewEvent(db, models.Event{})
	eventService := service.NewEventService(eventRepository)
	router.Post("/api/rest/v1/event", event.Add{Event: eventService}.ServeHTTP)
	router.Get("/api/rest/v1/event/{id}/{limit}", event.List{Event: eventService}.ServeHTTP)
	router.Put("/api/rest/v1/event/{id}", event.Edit{Event: eventService}.ServeHTTP)
	router.Get("/api/rest/v1/event/{id}", event.Get{Event: eventService}.ServeHTTP)
	router.Delete("/api/rest/v1/event/{id}", event.Delete{Event: eventService}.ServeHTTP)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
