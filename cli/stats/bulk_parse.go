package stats

import (
	"codecoach/types"
	"log"
	"os/exec"
	"strings"
)

// This is used to back populate all git history from a repository

func BulkParse() [][]types.Stats {
	log.Println("starting bulk load")
	cmd := exec.Command("git", "log", "--numstat", "-2")

	b, err := cmd.Output()

	if err != nil {
		log.Println("error")

		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	commitSignature := "commit "
	start := 0
	end := 0

	var commits [][]types.Stats
	for i, v := range lines {
		log.Println("line", v)
		if (strings.HasPrefix(v, commitSignature) && end != 0) || end == len(lines)-1 {
			log.Println("HAS PREFIX")

			loggedCommit := []byte(strings.Join(lines[start:end], "\n"))
			commits = append(commits, parseCommit(loggedCommit))
			end++
			start = i
		}
		end++
	}

	return commits

}
