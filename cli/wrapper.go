package main

import (
	"codecoach/commits"
	"os"
	"os/exec"

	"codecoach/cli/stats"
)

func main() {
	args := os.Args

	// bulk import flag
	if len(args) == 1 {
		stats.ReadGitLogs(commits.LogOptions{AllLogs: false})
		return
	}

	if args[1] == "bulk" {
		stats.ReadGitLogs(commits.LogOptions{AllLogs: true})
		return
	}
	// if len(args) > 2 {
	// 	if args[1] == "git" && args[2] == "commit" {
	// 		postCommandHook(args)
	// 	}
	// }

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
	stats.CollectCommitStats()
}
