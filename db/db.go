package db

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var Client *sql.DB

func init() {
	var dbPath string = getDatabasePath("./analytics.db")

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	Client = db
}

func getDatabasePath(relativePath string) string {
	absolutePath, err := filepath.Abs(relativePath)

	if err != nil {
		log.Fatal(err)
	}

	return absolutePath
}
