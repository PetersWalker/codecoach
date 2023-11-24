package handlers

import (
	"bytes"
	"codecoach/commits"
	"codecoach/utils"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestPostStat(t *testing.T) {
	//setup
	layout := "Mon Jan 02 15:04:05 2006 -0700"
	time, _ := time.Parse(layout, "Thu Nov 16 14:44:17 2023 -0500")
	json, _ := json.Marshal([]commits.Stats{{
		Filepath:        "cli/wrapper.go",
		LinesAdded:      12,
		LinesSubtracted: 51,
		Name:            "PetersWalker",
		Date:            time,
		CommitHash:      "c9fe1ef646916078b52540846e25b5a156e6eb39",
	}})

	//action
	body := bytes.NewBuffer(json)
	req, _ := http.NewRequest("POST", "http://localhost:8000/postStats", body)
	client := &http.Client{}
	res, _ := client.Do(req)
	//assert
	assert.DeepEqual(t, res.Status, "200 OK")
}

func TestCastToStats(t *testing.T) {
	//setup
	commit := commits.RawCommit{
		CommitHash: "2528f600f73947495c7396a0d6d5ff2f1a4d343c",
		Author:     "PetersWalker",
		Date:       "1700690643",
		Files: []commits.RawFile{
			{
				FilePath:   "cli/wrapper.go",
				Added:      "0",
				Subtracted: "1",
			},
			{
				FilePath:   "cli/wrap.go",
				Added:      "0",
				Subtracted: "2",
			},
		},
	}

	date := time.Unix(int64(1700690643), 0)

	desired := []commits.Stats{
		{
			Filepath:        "cli/wrapper.go",
			LinesAdded:      0,
			LinesSubtracted: 1,
			Name:            "PetersWalker",
			Date:            date,
			CommitHash:      "2528f600f73947495c7396a0d6d5ff2f1a4d343c",
		},
		{
			Filepath:        "cli/wrap.go",
			LinesAdded:      0,
			LinesSubtracted: 2,
			Name:            "PetersWalker",
			Date:            date,
			CommitHash:      "2528f600f73947495c7396a0d6d5ff2f1a4d343c",
		},
	}

	//act
	result := castToStats(commit)

	assert.DeepEqual(t, result, desired)
}

func TestNomalizDates(t *testing.T) {
	// setup
	commitStats := utils.CommitStatsDenormalizedExample()

	// act
	result := normalizeDates(commitStats, 3)
	desired := utils.CommitStatsNormalizedExample()

	//assert

	assert.DeepEqual(t, result, desired)

}
