package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"codecoach/cli/stats"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Println("codecoach: no arguments provided")
		return
	}

	// if I want
	if args[1] == "--bulk" {
		stats.BulkParse()
		return
	}

	output, _ := executeArgs(args)
	fmt.Printf("%v", string(output))

	if len(args) > 2 {
		if args[1] == "git" && args[2] == "commit" {
			postCommandHook(args)
		}
	}

	return
}

func executeArgs(args []string) ([]byte, error) {
	var command *exec.Cmd

	if len(args) == 2 {
		command = exec.Command(args[1])
	}

	if len(args) > 2 {
		command = exec.Command(args[1], args[2:]...)
	}

	output, err := command.Output()

	return output, err
}

func postCommandHook(args []string) {

	if args[1] == "git" && args[2] == "commit" {
		stats.CollectCommitStats()
	}
}
