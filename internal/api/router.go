package api

import (
	"encoding/json"
	"net/http"
	"time"
)

// Router mengelola rute HTTP dasar untuk inisiasi proyek.
type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	r := &Router{
		mux: http.NewServeMux(),
	}
	r.registerRoutes()
	return r
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Header CORS dasar untuk pengembangan
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	r.mux.ServeHTTP(w, req)
}

func (r *Router) registerRoutes() {
	// Endpoint dasar verifikasi sistem
	r.mux.HandleFunc("GET /healthz", r.handleHealth)
	
	// Catatan: Endpoint telemetri, histori, dan ekspor dataset CSV
	// akan diimplementasikan bertahap pada fase berikutnya.
}

func (r *Router) handleHealth(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"service":   "smart-pot-irrigation-backend",
		"status":    "healthy",
		"stage":     "project-initialization",
		"timestamp": time.Now().UTC(),
	})
}
