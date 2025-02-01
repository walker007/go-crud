package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-crud/api/dtos"
	"go-crud/api/helpers"
	"go-crud/api/models"
	"go-crud/database"
)

func NewHandler(db *database.DBApplication[*models.User]) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", handleListUsers(db))
			r.Post("/", handleCreateUser(db))
			r.Put("/{id:[0-9]+}", handleUpdateUser(db))
			r.Get("/{id:[0-9]+}", handleGetUserById(db))
			r.Delete("/{id:[0-9]+}", handleDeleteUser(db))
		})
	})

	return r
}
func handleListUsers(db *database.DBApplication[*models.User]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := db.GetAll()
		helpers.SendJson(w, dtos.Response{
			Data: users,
		}, http.StatusOK)
	}
}
func handleGetUserById(db *database.DBApplication[*models.User]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")
		id, _ := strconv.ParseUint(idString, 10, 64)
		if userData, err := db.GetById(id); err == nil {
			helpers.SendJson(w, dtos.Response{Data: userData}, http.StatusOK)
			return
		}
		helpers.SendJson(w, dtos.Response{Error: "Usuário Não encontrado"}, http.StatusNotFound)
	}
}
func handleCreateUser(db *database.DBApplication[*models.User]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			slog.Error("Não foi possível obter dados da requisição", "error", err)
			helpers.SendJson(w, dtos.Response{Error: "Payload Inválido"}, http.StatusUnprocessableEntity)
			return
		}
		helpers.SendJson(w, dtos.Response{Data: db.Insert(&user)}, http.StatusCreated)
		return
	}
}

func handleUpdateUser(db *database.DBApplication[*models.User]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")
		id, _ := strconv.ParseUint(idString, 10, 64)
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			slog.Error("Não foi possível obter dados da requisição", "error", err)
			helpers.SendJson(w, dtos.Response{Error: "Payload Inválido"}, http.StatusUnprocessableEntity)
			return
		}

		data, err := db.Update(id, &user)
		if err != nil {
			helpers.SendJson(w, dtos.Response{Error: "usuário Não encontrado"}, http.StatusNotFound)
			return
		}

		helpers.SendJson(w, dtos.Response{Data: data}, http.StatusOK)
	}
}

func handleDeleteUser(db *database.DBApplication[*models.User]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")
		id, _ := strconv.ParseUint(idString, 10, 64)
		err := db.Delete(id)
		if err != nil {
			helpers.SendJson(w, dtos.Response{Error: "usuário Não encontrado"}, http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
