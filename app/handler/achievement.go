package handler

import (
	"encoding/json"
	"net/http"
	"portfolio/app/repository"
)

type achievement struct {
	mysql repository.Mysql
}

type Achievement interface {
	GetAll(http.ResponseWriter, *http.Request)
}

var _ Achievement = &achievement{}

func NewAchievement(mysql repository.Mysql) Achievement {
	return &achievement{
		mysql: mysql,
	}
}

func (ah *achievement) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	achs, err := ah.mysql.GetAchievements()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(achs); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
