package main

import (
	"log"
	"net/http"
	"url-shortener/internal/config"
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sqlx.Connect("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer db.Close()

	repo := repository.NewURLRepository(db)
	svc := service.NewURLService(repo, cfg.BaseURL)
	h := handler.NewHandler(svc)

	r := chi.NewRouter()
	r.Post("/shorten", h.ShortenURL)
	r.Get("/{code}", h.ResolveURL)

	log.Printf("🚀 Server running on :%s\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal("Server failed:", err)
	}
}
