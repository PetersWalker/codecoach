package stats

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"codecoach/types"
)

func CollectCommitStats() {
	output, err := exec.Command("git", "diff", "HEAD").Output()

	if err != nil {
		log.Fatal("collectCommitStats", err)
	}
	stats := parseCommit(output)

	// possibly some input validation here on the git diff bytes?
	// e.g. if not valit log.Fatal("non standard git diff format %v", str(output))

	b, err := json.Marshal(stats)

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
		panic(err)
	}

	if res.StatusCode != 200 {
		panic(res.Status)
	}
}

func parseCommit(output []byte) types.Stats {
	str := strings.Split(string(output), "\n")
	for _, v := range str {
		fmt.Println(v)
	}
	diffLines := str[4]
	subtracted := diffLines[6:8]
	added := diffLines[12:14]

	return types.Stats{
		LinesAdded:      added,
		LinesSubtracted: subtracted,
	}
}
