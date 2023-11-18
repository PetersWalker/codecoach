package stats

import (
	"log"
	"os"
	"testing"
	"time"

	"gotest.tools/v3/assert"

	"codecoach/types"
)

var data = `commit c9fe1ef646916078b52540846e25b5a156e6eb39 (HEAD -> main)
Author: PetersWalker <petersinclairwalker@gmail.com>
Date:   Thu Nov 16 14:44:17 2023 -0500

    feat: postCommand in go routine

12      51      cli/wrapper.go
5       1       notes.md
`

func TestParseCommit(t *testing.T) {
	// setup
	bytes, err := os.ReadFile("./fixtures/git_log_numstat.txt")
	if err != nil {
		log.Fatal(err)
	}

	// act
	result := parseCommit(bytes)
	layout := "Mon Jan 02 15:04:05 2006 -0700"
	time, _ := time.Parse(layout, "Thu Nov 16 14:44:17 2023 -0500")
	var expected = []types.Stats{
		{
			Filepath:        "cli/wrapper.go",
			LinesAdded:      "12",
			LinesSubtracted: "51",
			Name:            "PetersWalker",
			Date:            time,
			CommitHash:      "c9fe1ef646916078b52540846e25b5a156e6eb39",
		},
		{
			Filepath:        "notes.md",
			LinesAdded:      "5",
			LinesSubtracted: "1",
			Name:            "PetersWalker",
			Date:            time,
			CommitHash:      "c9fe1ef646916078b52540846e25b5a156e6eb39",
		},
	}

	// assert
	AssertEqualSlices[types.Stats](t, result, expected)

}

func TestParseCommitSingleFileChange(t *testing.T) {
	// setup
	bytes, err := os.ReadFile("./fixtures/git_log_numstat_single_file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// act
	result := parseCommit(bytes)
	layout := "Mon Jan 02 15:04:05 2006 -0700"
	time, _ := time.Parse(layout, "Thu Nov 16 14:44:17 2023 -0500")
	var expected = []types.Stats{
		{
			Filepath:        "cli/wrapper.go",
			LinesAdded:      "12",
			LinesSubtracted: "51",
			Name:            "PetersWalker",
			Date:            time,
			CommitHash:      "c9fe1ef646916078b52540846e25b5a156e6eb39",
		},
	}

	// assert
	AssertEqualSlices[types.Stats](t, result, expected)

}

func TestParseCommitFileNameChange(t *testing.T) {
	// setup
	bytes, err := os.ReadFile("./fixtures/git_log_numstat_name_change.txt")
	if err != nil {
		log.Fatal(err)
	}

	// act
	result := parseCommit(bytes)
	layout := "Mon Jan 02 15:04:05 2006 -0700"
	time, _ := time.Parse(layout, "Thu Nov 16 14:44:17 2023 -0500")
	var expected = []types.Stats{
		{
			Filepath:        "{stats => cli/stats}/collect_commit_stats.go",
			LinesAdded:      "2",
			LinesSubtracted: "14",
			Name:            "PetersWalker",
			Date:            time,
			CommitHash:      "c9fe1ef646916078b52540846e25b5a156e6eb39",
		},
	}

	// assert
	AssertEqualSlices[types.Stats](t, result, expected)

}

func AssertEqualSlices[T comparable](t assert.TestingT, result []T, expected []T) {
	for i, v := range result {
		assert.DeepEqual(t, v, expected[i])
	}
}
