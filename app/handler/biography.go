package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	bios, err := bh.mysql.GetBiobraphies()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	if err = json.NewEncoder(w).Encode(bios); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
