package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Danuson17-8/corn-backend/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB(cfg *config.EnvConfig) {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQL")
}
