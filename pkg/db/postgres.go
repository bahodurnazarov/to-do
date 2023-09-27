package db

import (
	"database/sql"
	"fmt"
	"github.com/bahodurnazarov/to-do/pkg/models"
	_ "github.com/lib/pq"
	"log"
)

var (
	DB *sql.DB
)

func ConnectDB(cfg *models.Settings) (*sql.DB, error) {
	db, err := postgres(cfg)
	return db, err

}
func postgres(cfg *models.Settings) (*sql.DB, error) {
	database := cfg.Postgres

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		database.PG_host, database.PG_port, database.PG_user, database.PG_password, database.PG_name)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	log.Printf("*** POSTGRES *** -> success connect to Database by user %s", database.PG_user)

	DB = db

	return DB, nil
}

func CloseDB() error {
	return DB.Close()
}
