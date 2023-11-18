package stats

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"codecoach/types"
)

func CollectCommitStats() {
	output, err := exec.Command("git", "log", "--numstat", "-1").Output()

	if err != nil {
		log.Fatal("collectCommitStats", err)
	}
	stats := parseCommit(output)

	// possibly some input validation here on the git diff bytes?
	// e.g. if not valit log.Fatal("non standard git diff format %v", str(output))

	b, err := json.Marshal(stats)
	log.Println("json", b)

	if err != nil {
		log.Fatal(err)
	}

	body := bytes.NewBuffer(b)
	req, err := http.NewRequest("POST", "http://localhost:8000/postStats", body)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatal(res.Status)
	}

	log.Println("response", res)
}

func parseCommit(gitLogNumstat []byte) []types.Stats {
	str := strings.Split(string(gitLogNumstat), "\n")
	commitLine := str[0]
	nameLine := str[1]
	dateLine := str[2]
	diffLines := str[6:]

	commit := strings.Fields(commitLine)[1]
	name := strings.Fields(nameLine)[1]
	dateString := strings.Join(strings.Fields(dateLine)[1:], " ")
	layout := "Mon Jan 02 15:04:05 2006 -0700"
	date, _ := time.Parse(layout, dateString)

	var commitStats []types.Stats
	for _, line := range diffLines {

		// ignore empty lines typically at the end
		if len(line) == 0 {
			continue
		}

		diff := strings.Fields(line)
		added := diff[0]
		subtracted := diff[1]
		file := strings.Join(diff[2:], " ")
		commitStats = append(commitStats, types.Stats{
			Filepath:        file,
			LinesAdded:      added,
			LinesSubtracted: subtracted,
			Name:            name,
			CommitHash:      commit,
			Date:            date,
		})
	}
	return commitStats
}
