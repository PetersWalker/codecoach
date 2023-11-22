package handlers

import (
	"bytes"
	"codecoach/db"
	"codecoach/types"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestSaveStats(t *testing.T) {

	// setup
	layout := "Mon Jan 02 15:04:05 2006 -0700"
	time, _ := time.Parse(layout, "Thu Nov 16 14:44:17 2023 -0500")
	stats := []types.Stats{
		{
			Filepath:        "cli/wrapper.go",
			LinesAdded:      "12",
			LinesSubtracted: "51",
			Name:            "PetersWalker",
			Date:            time,
			CommitHash:      "c9fe1ef646916078b52540846e25b5a156e6eb39",
		},
	}

	// action
	result, err := recordStatsData(db.Client, stats)
	assert.NilError(t, err)
	assert.DeepEqual(t, result, stats)
}

func TestPostStat(t *testing.T) {
	//setup
	layout := "Mon Jan 02 15:04:05 2006 -0700"
	time, _ := time.Parse(layout, "Thu Nov 16 14:44:17 2023 -0500")
	json, _ := json.Marshal([]types.Stats{{
		Filepath:        "cli/wrapper.go",
		LinesAdded:      "12",
		LinesSubtracted: "51",
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
