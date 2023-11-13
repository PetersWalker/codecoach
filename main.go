package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/printStats", printStatsHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func printStatsHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	err:= json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, data)
}
