package handler

import (
	"awesomeProject-URL/internal/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (h *Handler) createUrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var url model.URL

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	shortUrl, err := h.service.CreateUrl(url)
	if err != nil {
		log.Printf("Create url error %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		log.Printf("[H] Short url received %s\n", shortUrl)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(shortUrl)
	}
}

func (h *Handler) getShortUrl(w http.ResponseWriter, r *http.Request) {

	shortUrl, _ := mux.Vars(r)["short_url"]

	url, err := h.service.GetUrl(shortUrl)
	if err != nil {
		log.Printf("Get url error %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		log.Printf("[H] Url received %s\n", url)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(url)
	}
}
