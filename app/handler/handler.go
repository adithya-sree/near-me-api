package handler

import (
	"encoding/json"
	"fmt"
	"nearme-api/app/db"
	"net/http"
)

type Handler struct {
	*db.Client
}

type Response struct {
	Message string `json:"message"`
}

func NewHandler(db *db.Client) *Handler {
	return &Handler{db}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func checkforHeader(r *http.Request, headerKey string) (string, error) {
	header := r.Header.Get(headerKey)
	if header == "" {
		return "", fmt.Errorf("required header " + headerKey + " was not in the request")
	}
	return header, nil
}
