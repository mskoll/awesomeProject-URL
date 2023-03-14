package handler

import (
	"awesomeProject-URL/internal/service"
	"fmt"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", h.createUrl).Methods("POST")
	router.HandleFunc("/{short_url}", h.getUrl).Methods("GET")
	fmt.Printf("ROUTES INITS\n")
	return router
}
