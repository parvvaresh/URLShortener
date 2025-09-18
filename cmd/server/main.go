package main

import (
	"log"
	"net/http"
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func main() {
	// اتصال به دیتابیس
	db, err := sqlx.Connect("postgres", "postgres://user:password@localhost:5432/urlshortener?sslmode=disable")
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// لایه‌ها
	repo := repository.NewURLRepository(db)
	svc := service.NewURLService(repo)
	h := handler.NewHandler(svc)

	// روتینگ
	r := chi.NewRouter()
	r.Post("/shorten", h.ShortenURL)
	r.Get("/{code}", h.ResolveURL)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
