package postgres

import (
	"fmt"
	"log"

	config "example.com/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func GetConnection() *sqlx.DB {
	if db != nil {
		return db
	}

	cfg := config.LoadConfig().DB

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Name)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed database connection!")
		return nil
	}

	return db
}
