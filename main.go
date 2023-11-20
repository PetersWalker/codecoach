package main

import (
	"codecoach/handlers"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	counter := 0
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, counter)
		counter = counter + 1
	}

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/home", h1)
	http.HandleFunc("/health", healthhandler)
	http.HandleFunc("/postStats", handlers.PostStatsHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	fmt.Println("listening on local:8000")

}

// handler echoes the Path component of the request URL r.
func healthhandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "healthy")
}
