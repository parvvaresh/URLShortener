package handler

import (
	"encoding/json"
	"net/http"
	"url-shortener/internal/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	svc **service.URLService
}
