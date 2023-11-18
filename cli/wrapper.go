package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"codecoach/stats"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		log.Print("no arguments")
		return
	}

	executeArgs(args)
	go postCommandHook(args)
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

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", output)
}

func postCommandHook(args []string) {

	if args[1] == "git" && args[2] == "commit" {
		stats.CollectCommitStats()
	}
}
