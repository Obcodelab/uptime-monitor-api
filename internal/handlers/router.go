// Package handlers holds the Chi router and the HTTP handlers.
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewRouter builds the Chi router. Chi's mux implements http.Handler, so it
// plugs straight into http.Server; routes and middleware are added with
// Get/Post/... and Use (later steps will group routes with r.Route/r.Group).
func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/healthz", healthz)
	return r
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
