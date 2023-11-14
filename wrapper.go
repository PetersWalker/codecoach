package main

import (
	"log"
	"os/exec"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		command := os.Args[1]

		log.Printf("Running %v ", command)
		output, err := exec.Command(command).Output()

		if err != nil{
			log.Fatal(err)
		}

		fmt.Printf("%s", output)
	}

	if len(os.Args) > 2 {
		command := os.Args[1]
		// args := ""
		// for _, v := range os.Args[2:] {
		// 	args = args + v + " "
		// }
		output, err := exec.Command(command, os.Args[2:]...).Output()

		if err != nil{
			log.Fatal(err)
		}

		fmt.Printf("%s", output)
	}
}
