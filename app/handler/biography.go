package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"portfolio/app/repository"
)

type biography struct {
	mysql repository.Mysql
}

type Biography interface {
	GetAll(http.ResponseWriter, *http.Request)
}

var _ Biography = &biography{}

func NewBiography(mysql repository.Mysql) Biography {
	return &biography{
		mysql: mysql,
	}
}

func (bh *biography) GetAll(w http.ResponseWriter, r *http.Request) {
	allowedDomain := os.Getenv("ALLOWED_DOMAIN")

	w.Header().Set("Access-Control-Allow-Origin", allowedDomain)
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
