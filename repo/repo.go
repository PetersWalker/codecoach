package repo

import (
	"codecoach/commits"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

func RecordStatsData(db *sql.DB, data []commits.Stats) ([]commits.Stats, error) {
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
			commit.Date.Format("2006-01-02"),
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

func GetCommitData(db *sql.DB) (any, error) {
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
