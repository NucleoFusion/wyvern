package pg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectPG() (*sql.DB, error) {
	host := os.Getenv("POSTGRES_HOST") // e.g. "localhost" or "postgres" in docker
	port := 5432                       // Default Postgres port
	user := os.Getenv("POSTGRES_USER") // e.g. "postgres"
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("[CONNECTED] Postgres")
	return db, nil
}
