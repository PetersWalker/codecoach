package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"codecoach/db"
	"codecoach/types"
)

func Healthhandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "healthy")
}

func PostStatsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("postStats called")
	var data []types.Stats
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Fprint(w, err)
	}

	log.Println(data)
	saved_data, err := saveStatsData(db.Client, data)

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
	}
	log.Println(saved_data)
}

func saveStatsData(db *sql.DB, data []types.Stats) ([]types.Stats, error) {
	var err error

	for _, commit := range data {
		query := `
		insert into 
		commit_stats (lines_added, lines_subtracted, name, file_path, date, commit_hash)
		values (?, ?, ?, ?, ?, ?)`

		_, err := db.Query(query,
			commit.LinesAdded,
			commit.LinesSubtracted,
			commit.Name,
			commit.Filepath,
			commit.Date,
			commit.CommitHash,
		)

		if err != nil {
			err = fmt.Errorf("saveStatsData: failed executing db query: %w", err)
			log.Println(err)
			break
		}
		log.Printf("saveStatsData rows:%+v \n", data)

	}

	return data, err

}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	commitData, _ := getCommitData(db.Client)
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, commitData)
}

func getCommitData(db *sql.DB) (any, error) {
	query := `
	select 
		sum(lines_added), 
		sum(lines_subtracted), 
		commit_hash, 
		date 
	from commit_stats 
	group by date;
	`

	rows, err := db.Query(query)

	var added int
	var subtracted int
	var hash string
	var date string

	type Counts struct {
		Added      int
		Subtracted int
		Hash       string
		Date       string
	}

	var data []Counts

	for rows.Next() {
		rows.Scan(&added, &subtracted, &hash, &date)
		data = append(
			data,
			Counts{
				Added:      added,
				Subtracted: subtracted,
				Hash:       hash,
				Date:       date,
			})
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := json.Marshal(data)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return string(result), err
}
