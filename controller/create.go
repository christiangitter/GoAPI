package controller

import (
	"GoAPI/model"
	"net/http"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// take the Data
			// save it to the database
			if err := model.Create(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
