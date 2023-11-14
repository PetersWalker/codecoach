package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io"
)

func main() {
	http.HandleFunc("/", healthhandler) // each request calls handler
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
}
