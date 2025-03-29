package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wwwstr/internal/db/model"

	"gorm.io/gorm"
)

func QuestionsSetHandlerW(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		questionID := r.Form.Get("id")
		intQuestionID, err := strconv.Atoi(questionID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response := r.Form.Get("response")
		question, err := model.AddQuestion(db, intQuestionID, response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(question)
	}
}
