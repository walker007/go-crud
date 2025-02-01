package helpers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"go-crud/api/dtos"
)

func SendJson(w http.ResponseWriter, resp dtos.Response, status int) {
	w.Header().Set("Accept", "application/json")

	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err)
		SendJson(w, dtos.Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}
}
