package repo

import (
	"codecoach/commits"
	"database/sql"
	"fmt"
	"log"
	"time"
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

func GetCommitData(db *sql.DB) ([]commits.Stats, error) {
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

	var data []commits.Stats

	for rows.Next() {
		var s commits.Stats
		var stringDate string
		err := rows.Scan(&s.LinesAdded, &s.LinesSubtracted, &s.CommitHash, &stringDate)

		if err != nil {
			panic(err)
		}

		date, err := time.Parse(time.DateOnly, stringDate)
		if err != nil {
			panic(err)
		}
		s.Date = date

		if err = rows.Err(); err != nil {
			panic(err) // Error related to the iteration of rows
		}

		data = append(data, s)
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, err
}
