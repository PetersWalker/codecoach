package repo

import (
	"codecoach/commits"
	"codecoach/db"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestRecordStats(t *testing.T) {

	// setup
	layout := "Mon Jan 02 15:04:05 2006 -0700"
	time, _ := time.Parse(layout, "Thu Nov 16 14:44:17 2023 -0500")
	stats := []commits.Stats{
		{
			Filepath:        "cli/wrapper.go",
			LinesAdded:      12,
			LinesSubtracted: 51,
			Name:            "PetersWalker",
			Date:            time,
			CommitHash:      "c9fe1ef646916078b52540846e25b5a156e6eb39",
		},
	}

	// action
	result, err := RecordStatsData(db.Client, stats)
	assert.NilError(t, err)
	assert.DeepEqual(t, result, stats)
}
