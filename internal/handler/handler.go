package handler

import (
	"encoding/json"
	"net/http"
	"url-shortener/internal/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	svc *service.URLService
}

func NewHandler(svc *service.URLService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		URL string `json:"url"`
	}
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	shortCode, err := h.svc.Shorten(req.URL)
	if err != nil {
		http.Error(w, "failed to shorten url", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"short_url": "http://localhost:8080/" + shortCode,
	})
}

func (h *Handler) ResolveURL(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	url, err := h.svc.Resolve(code)
	if err != nil {
		http.Error(w, "url not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
