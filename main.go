package main

import (
	"codecoach/handlers"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/home", handlers.HandleHome)
	http.HandleFunc("/health", handlers.Healthhandler)
	http.HandleFunc("/postStats", handlers.PostStatsHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	fmt.Println("listening on local:8000")

}
