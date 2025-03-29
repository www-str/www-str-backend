package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wwwstr/internal/db/model"

	"gorm.io/gorm"
)

func RatingsSetHandlerW(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link := r.Form.Get("link")
		rate := r.Form.Get("rate")
		intRate, err := strconv.Atoi(rate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rating, err := model.AddRating(db, link, intRate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(rating)
	}
}
