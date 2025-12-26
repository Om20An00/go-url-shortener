package storage

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Exported context for queries
var Ctx = context.Background()

var DB *sql.DB

func InitPostgres() {
	dsn := os.Getenv("POSTGRES_DSN")

	// fallback for local (non-docker) runs
	if dsn == "" {
		dsn = "postgres://postgres:Omanand%402000@localhost:5432/urlshortener?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Postgres open error:", err)
	}

	if err = db.PingContext(Ctx); err != nil {
		log.Fatal("Postgres ping error:", err)
	}

	DB = db
	log.Println("Connected to PostgreSQL")
}
