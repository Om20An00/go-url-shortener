package storage

import (
	"database/sql"
	"log"
)

// InsertURL saves the short URL mapping in Postgres
func InsertURL(short, original string) error {
	_, err := DB.Exec("INSERT INTO urls (short, original) VALUES ($1, $2)", short, original)
	if err != nil {
		log.Println("InsertURL error:", err)
	}
	return err
}

// GetOriginalURL fetches the original URL from Postgres
func GetOriginalURL(short string) (string, error) {
	var original string
	err := DB.QueryRow("SELECT original FROM urls WHERE short = $1", short).Scan(&original)
	if err == sql.ErrNoRows {
		return "", err
	} else if err != nil {
		log.Println("GetOriginalURL error:", err)
		return "", err
	}
	return original, nil
}
