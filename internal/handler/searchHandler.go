package handler

import (
	"net/http"

	"resty.dev/v3"
)

func SearchGetHandlerW(googleApiKey, searchKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query == "" {
			http.Error(w, "Missing 'q' parameter", http.StatusBadRequest)
			return
		}
		client := resty.New()
		defer client.Close()
		data, err := client.R().
			SetQueryParam("key", googleApiKey).
			SetQueryParam("cx", searchKey).
			SetQueryParam("safe", "active").
			SetQueryParam("lr", "lang_ru").
			SetQueryParam("gl", "ru").
			SetQueryParam("q", query).
			Get("https://www.googleapis.com/customsearch/v1")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(data.Bytes())
	}
}
