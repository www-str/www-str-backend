package db

import (
	"fmt"
	"log"
	"wwwstr/internal/db/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
	Database string
}

func migrateAll(db *gorm.DB) {
	db.AutoMigrate(
		model.Rating{},
		model.Question{},
	)
}

func (conf *DbConfig) InitConnection() (*gorm.DB, error) {
	// dsn := "host=" + conf.Host +
	// 	" port=" + conf.Port +
	// 	" user=" + conf.UserName +
	// 	" password=" + conf.Password +
	// 	" dbname=" + conf.Database +
	// 	" sslmode=disable"
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		conf.UserName,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
		return nil, err
	}

	migrateAll(db)
	return db, nil
}
