package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io"
	"database/sql"
	"path/filepath"

	_ "modernc.org/sqlite"

	_ "codecoach/types"
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
}

// handler echoes the Path component of the request URL r.
func healthhandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "healthy")
}

func postStatsHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	err:= json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Printf("%v\n", data)
	query:= `insert into commit_stats values(?, ?)`
	db.Query(query, data["LinesAdded"], data["LinesSubtracted"])
}

func getDatabasePath(relativePath string) string{
	absolutePath, err := filepath.Abs(relativePath)

	if err != nil {
		log.Fatal(err)
	}

	return absolutePath
}

func insertStats(stats string) string {
	return stats
}
