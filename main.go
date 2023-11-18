package main

import (
	"codecoach/handlers"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	http.HandleFunc("/", healthhandler)
	http.HandleFunc("/postStats", handlers.PostStatsHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	fmt.Println("listening on local:8000")
}

// handler echoes the Path component of the request URL r.
func healthhandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "healthy")
}
