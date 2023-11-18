package stats

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	fmt.Printf("%s", diffLines)
	var commitStats []types.Stats

	commit := strings.Fields(commitLine)[1]
	name := strings.Fields(nameLine)[1]
	dateString := strings.Join(strings.Fields(dateLine)[1:], " ")
	layout := "Mon Jan 02 15:04:05 2006 -0700"
	date, _ := time.Parse(layout, dateString)
	fmt.Println("date", date)

	fmt.Println("diffLines", diffLines)

	for _, line := range diffLines {
		diff := strings.Fields(line)
		fmt.Println("diff", diff)

		added := diff[0]
		subtracted := diff[1]
		file := diff[2]
		commitStats = append(commitStats, types.Stats{
			Filepath:        file,
			LinesAdded:      added,
			LinesSubtracted: subtracted,
			Name:            name,
			Commit:          commit,
			Date:            date,
		})
	}

	fmt.Println("commit stats", commitStats)

	return commitStats
}
