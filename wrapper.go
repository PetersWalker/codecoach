package main

import (
	"log"
	"os/exec"
	"fmt"
	"os"
)

func main() {
	executeArgs(os.Args)
}

func executeArgs(args []string) {
	if len(args) == 1 {
		return
	}

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
