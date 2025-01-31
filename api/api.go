package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go-crud/api/models"
	"go-crud/database"
)

func NewHandler(db *database.DBApplication[*models.User]) http.Handler {

	r := chi.NewMux()

	return r
}
