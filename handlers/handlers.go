package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"codecoach/commits"
	"codecoach/db"
	"codecoach/repo"
	"codecoach/utils"
)

func Healthhandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "healthy")
}

func PostStatsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("postStats called")
	var data []commits.Stats
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Fprint(w, err)
	}

	log.Println(data)
	saved_data, err := repo.RecordStatsData(db.Client, data)

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
	}
	log.Println(saved_data)
}

func PostCommitsBulk(w http.ResponseWriter, r *http.Request) {
	start := time.Now().UnixMilli()
	log.Println("postCommitsBulk")
	var commits []commits.RawCommit

	err := json.NewDecoder(r.Body).Decode(&commits)

	if err != nil {
		fmt.Fprint(w, err)
	}

	for _, commit := range commits {
		stats := castToStats(commit)
		_, err := repo.RecordStatsData(db.Client, stats)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
	}

	end := time.Now().UnixMilli()
	log.Println("completed postCommits in: ", end-start, "ms")
}

func castToStats(commit commits.RawCommit) []commits.Stats {
	var stats []commits.Stats
	for _, file := range commit.Files {
		unixInt, err := strconv.ParseInt(commit.Date, 10, 64)

		if err != nil {
			log.Panic("castToStats: date string conversion: ", commit.Date, err)
		}

		date := time.Unix(unixInt, 0)

		// todo fix messy validation
		if file.Added == "-" {
			file.Added = "0"
		}

		added, err := strconv.Atoi(file.Added)

		if err != nil {
			log.Panic("castToStats: added string conversion: ", file.Added, err)
		}

		// todo fix messy validation
		if file.Subtracted == "-" {
			file.Subtracted = "0"
		}

		subtracted, err := strconv.Atoi(file.Subtracted)

		if err != nil {
			log.Panic("castToStats: subtracted string conversion: ", file.Subtracted, err)
		}

		stat := commits.Stats{
			Filepath:        file.FilePath,
			LinesAdded:      added,
			LinesSubtracted: subtracted,
			Name:            commit.Author,
			Date:            date,
			CommitHash:      commit.CommitHash,
		}

		stats = append(stats, stat)
	}

	return stats

}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	//queryvars =  validator.validate(, )
	days := r.URL.Query().Get("days")
	commitStats, _ := repo.GetAllCommitData(db.Client)

	// todo proper validation
	if days == "" {
		days = "7"
	}

	numberofDays, err := strconv.Atoi(days)
	if err != nil {
		log.Println("HandleHome: atoi conversion failed", err)
		log.Println("test\n test")
		w.Write([]byte("invalid parameter for query string: days"))
		log.Panic(err)
		return
	}

	var firstDate = time.Now().AddDate(0, 0, -numberofDays)
	normalizedCommitStats := normalizeDates(commitStats, numberofDays, firstDate)

	result, err := json.Marshal(normalizedCommitStats)

	if err != nil {
		log.Println(err)
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, string(result))
}

func normalizeDates(commitStats []commits.Stats, days int, firstDate time.Time) []commits.Stats {
	var zeros []commits.Stats
	var d = 0

	for d < days {
		zeros = append(zeros, utils.CommitStatNil(firstDate.AddDate(0, 0, d)))
		d++
	}

	statsIndex := 0

	for i, nilCommit := range zeros {
		if statsIndex == len(commitStats) {
			break
		}

		nilDate := nilCommit.Date
		commitDate := commitStats[statsIndex].Date

		if (nilDate.Day() == commitDate.Day()) && (nilDate.Month() == commitDate.Month()) && (nilDate.Year() == commitDate.Year()) {
			zeros[i] = commitStats[statsIndex]
			statsIndex++
		}
	}

	return zeros
}

func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	//queryvars =  validator.validate(, )
	window := r.URL.Query().Get("window")
	var numberOfDays int

	if window == "week" {
		numberOfDays = 7
	}

	if window == "month" {
		numberOfDays = 30
	}

	if window == "year" {
		numberOfDays = 365
	}

	if window == "" {
		numberOfDays = 7
	}

	var firstDate = time.Now().AddDate(0, 0, -numberOfDays)
	commitStats, _ := repo.GetCommitDataAfterDate(db.Client, firstDate)
	normalizedCommitStats := normalizeDates(commitStats, numberOfDays, firstDate)
	result, err := json.Marshal(normalizedCommitStats)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	io.WriteString(w, string(result))
}
