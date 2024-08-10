package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sakarbaral/personalized-recommendation-engine/services/recommendation-engine/pkg/config"
)

var db *sql.DB

func GetDB() *sql.DB {
	if db != nil {
		return db
	}
	var err error
	db, err = sql.Open("mysql", config.AppConfig.SQLConfig.ConnectionURL)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}
