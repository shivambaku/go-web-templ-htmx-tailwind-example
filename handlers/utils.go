package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/shivambaku/go-web-templ-htmx-tailwind-demo/internal/database"
)

type Handler struct {
	DB *database.Queries
}

func response(w http.ResponseWriter, code int, data []byte) {
	w.WriteHeader(code)
	_, err := w.Write(data)
	if err != nil {
		log.Printf("Error writing response: %s", err)
	}
}

func responseJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response(w, code, data)
}

func responseError(w http.ResponseWriter, code int, message string) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	if code > 499 {
		log.Printf("Responding with 5XX error: %s", message)
		message = "Internal server error"
	}
	responseJSON(w, code, errorResponse{Error: message})
}

func responseView(w http.ResponseWriter, r *http.Request, component templ.Component) {
	component.Render(r.Context(), w)
}
