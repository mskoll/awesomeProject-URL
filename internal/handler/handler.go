package handler

import (
	"awesomeProject-URL/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// InitRoutes инициализация эндпойнтов
func (h *Handler) InitRoutes() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", h.createUrl).Methods("POST")
	router.HandleFunc("/{short_url}", h.getShortUrl).Methods("GET")
	return router
}
