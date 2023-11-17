package stats

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var data = `commit c9fe1ef646916078b52540846e25b5a156e6eb39 (HEAD -> main)
Author: PetersWalker <petersinclairwalker@gmail.com>
Date:   Thu Nov 16 14:44:17 2023 -0500

    feat: postCommand in go routine

12      51      cli/wrapper.go
5       1       notes.md
`

func TestParseCommit(t *testing.T) {
	// get fixture
	bytes, err := os.ReadFile("./fixtures/git_log_numstat.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bytes)
}
