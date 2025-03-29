package wwwstr

import (
	"net/http"
	"wwwstr/internal/db"
	"wwwstr/internal/handler"

	"github.com/gorilla/mux"
	"github.com/pgmod/envconfig"
)

func Listen() error {
	err := envconfig.Load()
	if err != nil {
		return err
	}
	dbConf := db.DbConfig{
		Host:     envconfig.Get("DB_HOST", "localhost"),
		Port:     envconfig.Get("DB_PORT", "5432"),
		UserName: envconfig.Get("DB_USERNAME", "login"),
		Password: envconfig.Get("DB_PASSWORD", "password"),
		Database: envconfig.Get("DB_NAME", "postgres"),
	}
	db, err := dbConf.InitConnection()
	if err != nil {
		return err
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/ratings/set", handler.RatingsSetHandlerW(db)).Methods("POST")
	router.HandleFunc("/api/questions/set", handler.QuestionsSetHandlerW(db)).Methods("POST")
	return http.ListenAndServe(":"+envconfig.Get("HTTP_PORT", "8090"), router)
}
