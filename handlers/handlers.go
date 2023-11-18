package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"codecoach/db"
	"codecoach/types"
)

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
