package main

import (
	"log"
	"os/exec"
	"fmt"
	"os"
	"strings"
	// "bytes"
)

func main() {
	args := os.Args

	
	if len(args) < 3 {
		return
	}

	if args[1] == "git" && args[2] == "commit" {
		collectCommitStats()
		return
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
	output, err := exec.Command("git", "diff").Output()

	if err != nil {
		log.Fatal("collectCommiStats", err)
	}
	str := strings.Split(string(output), "\n")

	// possibly some input validation here on the git diff bytes?
	// e.g. if not valit log.Fatal("non standard git diff format %v", str(output))
	diffLines := str[4]
	subtracted := diffLines[6:8]
	added := diffLines[12:14]

	fmt.Printf("%v, %v", added, subtracted)
}