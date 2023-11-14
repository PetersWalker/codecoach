package main

import (
	"log"
	"os/exec"
	"fmt"
	"os"
	"net/http"
	"bytes"
	"strings"
	"encoding/json"
)

func main() {
	args := os.Args

	
	if len(args) < 3 {
		return
	}

	if args[1] == "git" && args[2] == "commit" {
		collectCommitStats()
	}
	
	executeArgs(args)
}

func executeArgs(args []string) {

	var command *exec.Cmd

	if len(args) == 2 {
		command = exec.Command(args[1])
	}

	if len(args) > 2 {
		command = exec.Command(args[1], args[2:]...)
	}

	output, err := command.Output()
	if err != nil{
			log.Fatal(err)
		}

	fmt.Printf("%s", output)
}

func collectCommitStats() {
	output, err := exec.Command("git", "diff", "HEAD").Output()

	if err != nil {
		log.Fatal("collectCommitStats", err)
	}
	str := strings.Split(string(output), "\n")

	// possibly some input validation here on the git diff bytes?
	// e.g. if not valit log.Fatal("non standard git diff format %v", str(output))
	for _, v := range str {
		fmt.Println(v)
	}
	diffLines := str[4]
	subtracted := diffLines[6:8]
	added := diffLines[12:14]

	type Stats struct {
		LinesAdded string
		LinesSubtracted string
	}

	stats := Stats{
		LinesAdded: added,
		LinesSubtracted: subtracted,
	}

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