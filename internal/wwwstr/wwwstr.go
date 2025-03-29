package wwwstr

import (
	"net/http"
	"wwwstr/internal/db"
	"wwwstr/internal/handler"

	"github.com/gorilla/mux"
	"github.com/pgmod/envconfig"
)

func Listen() error {
	dbConf := db.DbConfig{
		Host:     envconfig.Get("POSTGRES_HOST", "0.0.0.0"),
		Port:     envconfig.Get("POSTGRES_PORT", "5432"),
		UserName: envconfig.Get("DB_USER", "postgres"),
		Password: envconfig.Get("POSTGRES_PASSWORD", "pg_password"),
		Database: envconfig.Get("DB_NAME", "back_db"),
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
