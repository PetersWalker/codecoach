package utils

import (
	"codecoach/commits"
	"time"
)

var today = time.Now()
var yesterday = today.AddDate(0, 0, -1)
var twoDaysAgo = today.AddDate(0, 0, -2)

func CommitStatsDenormalizedExample() []commits.Stats {
	return []commits.Stats{
		{
			Filepath:        "cli/wrap.go",
			LinesAdded:      0,
			LinesSubtracted: 2,
			Name:            "PetersWalker",
			Date:            twoDaysAgo,
			CommitHash:      "2528f600f73947495c7396a0d6d5ff2f1a4d343c",
		},
		{
			Filepath:        "cli/wrapper.go",
			LinesAdded:      0,
			LinesSubtracted: 1,
			Name:            "PetersWalker",
			Date:            today,
			CommitHash:      "2528f600f73947495c7396a0d6d5ff2f1a4d343c",
		},
	}
}

func CommitStatsNormalizedExample() []commits.Stats {
	return []commits.Stats{
		{
			Filepath:        "cli/wrap.go",
			LinesAdded:      0,
			LinesSubtracted: 2,
			Name:            "PetersWalker",
			Date:            twoDaysAgo,
			CommitHash:      "2528f600f73947495c7396a0d6d5ff2f1a4d343c",
		},
		{
			Filepath:        "",
			LinesAdded:      0,
			LinesSubtracted: 0,
			Name:            "",
			Date:            yesterday,
			CommitHash:      "",
		},
		{
			Filepath:        "cli/wrapper.go",
			LinesAdded:      0,
			LinesSubtracted: 1,
			Name:            "PetersWalker",
			Date:            today,
			CommitHash:      "2528f600f73947495c7396a0d6d5ff2f1a4d343c",
		},
	}
}

func CommitStatNil(date time.Time) commits.Stats {
	return commits.Stats{
		Filepath:        "",
		LinesAdded:      0,
		LinesSubtracted: 0,
		Name:            "",
		Date:            date,
		CommitHash:      "",
	}
}
