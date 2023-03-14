package handler

import (
	"awesomeProject-URL/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) createUrl(w http.ResponseWriter, r *http.Request) {
	log.Printf("REQUEST\n")
	w.Header().Set("Content-Type", "application/json")
	var url model.URL

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	url.ShortUrl = url.Url
	id, err := h.service.CreateUrl(url)
	if err != nil {

		log.Printf("create url error %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
	fmt.Printf("URL created %d", id)
}

func (h *Handler) getUrl(w http.ResponseWriter, r *http.Request) {

}
