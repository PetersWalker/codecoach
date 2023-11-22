package stats

import (
	"bufio"
	"bytes"
	"codecoach/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"sync"
)

func TokenizeGitLogs(options LogOptions) []RawCommit {
	commands := []string{"git", "log", "--numstat", "--date=raw"}

	if options.AllLogs == false {
		commands = append(commands, "-1")
	}

	cmd := exec.Command(commands[0], commands[1:]...)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()

	scanner := bufio.NewScanner(stdout)

	var commit RawCommit
	var commitList []RawCommit
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)

		const commitSignature string = "commit "
		const authorSignature string = "Author: "
		const dateSignature string = "Date: "
		const messageSignature string = "    "

		if strings.HasPrefix(s, commitSignature) {
			// if we reach a new commit, append the previously collected commit
			if commit.CommitHash != "" {
				commitList = append(commitList, commit)
				commit = RawCommit{}
			}

			commit.CommitHash = strings.Fields(s)[1]
			continue
		}

		if strings.HasPrefix(s, authorSignature) {
			commit.Author = strings.Fields(s)[1]
			continue
		}

		if strings.HasPrefix(s, dateSignature) {
			commit.Date = strings.Join(strings.Fields(s)[1:], " ")
			continue
		}

		if strings.HasPrefix(s, messageSignature) {
			continue
		}

		if len(s) == 0 {
			continue
		}

		rawFile := parseFileChangeLine(s)
		commit.Files = append(commit.Files, rawFile)

		if err := scanner.Err(); err != nil {
			if err != io.EOF {
				fmt.Println("os.Stderr, err")
			}
		}
	}

	// grab the last commit after EOF
	commitList = append(commitList, commit)

	fmt.Printf("%+v\n", commitList)
	return commitList
}

func FlushCommits(rawCommits []RawCommit) {
	var wg sync.WaitGroup
	client := &http.Client{}
	wg.Add(1)
	go BulkPostCommit(client, rawCommits)
	wg.Wait()
}

func BulkPostCommit(client *http.Client, commits []RawCommit) {
	b, err := json.Marshal(commits)
	log.Println("json", b)

	if err != nil {
		log.Fatal(err)
	}

	body := bytes.NewBuffer(b)

	req, err := http.NewRequest("POST", "http://localhost:8000/postStats/bulk", body)

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatal(res.Status)
	}

	log.Println(res.StatusCode)
}

func PostStats(stats []types.Stats) {
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
}
