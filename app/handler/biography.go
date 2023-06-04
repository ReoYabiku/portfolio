package handler

import (
	"encoding/json"
	"net/http"
	"portfolio/app/repository"
)

type biographyHandler struct {
	mysql repository.Mysql
}

type BiographyHandler interface {
	GetAll(http.ResponseWriter, *http.Request)
}

var _ BiographyHandler = &biographyHandler{}

func NewBiographyHandler(mysql repository.Mysql) BiographyHandler {
	return &biographyHandler{
		mysql: mysql,
	}
}

func (bh *biographyHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Origin, X-Csrftoken, Accept, Cookie")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	bios, err := bh.mysql.GetBiobraphies()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(bios); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
