package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"go-crud/api"
	"go-crud/api/models"
	"go-crud/database"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		os.Exit(1)
	}
	slog.Info("the Service is Down")
}

func run() error {
	db := database.NewApplication[*models.User]()
	handler := api.NewHandler(db)
	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil

}
