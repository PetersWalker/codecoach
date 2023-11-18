package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func main() {
	dbPath := getDatabasePath("./analytics.db")
	var err error
	db, err = sql.Open("sqlite", dbPath)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", healthhandler)
	http.HandleFunc("/postStats", postStatsHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	fmt.Println("listening on local:8000")
}

// handler echoes the Path component of the request URL r.
func healthhandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "healthy")
}

func postStatsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("postStats called")
	var data []map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Fprint(w, err)
	}
	log.Println(data)
	for _, commit := range data {
		query := `insert into commit_stats values(?, ?)`

		db.Query(query, commit["LinesAdded"], commit["LinesSubtracted"])
	}

}

func getDatabasePath(relativePath string) string {
	absolutePath, err := filepath.Abs(relativePath)

	if err != nil {
		log.Fatal(err)
	}

	return absolutePath
}

func insertStats(stats string) string {
	return stats
}
