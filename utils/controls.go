package utils

import (
	"codecoach/commits"
	"time"
)

func CommitStatsDenormalizedExample(firstDate time.Time) []commits.Stats {

	var twoFromNow = firstDate.AddDate(0, 0, 2)
	return []commits.Stats{
		{
			Filepath:        "cli/wrap.go",
			LinesAdded:      0,
			LinesSubtracted: 2,
			Name:            "PetersWalker",
			Date:            firstDate,
			CommitHash:      "2528f600f73947495c7396a0d6d5ff2f1a4d343c",
		},
		{
			Filepath:        "cli/wrapper.go",
			LinesAdded:      0,
			LinesSubtracted: 1,
			Name:            "PetersWalker",
			Date:            twoFromNow,
			CommitHash:      "2528f600f73947495c7396a0d6d5ff2f1a4d343c",
		},
	}
}

func CommitStatsNormalizedExample(firstDate time.Time) []commits.Stats {
	var tomorrow = firstDate.AddDate(0, 0, 1)
	var twoFromNow = firstDate.AddDate(0, 0, 2)

	return []commits.Stats{
		{
			Filepath:        "cli/wrap.go",
			LinesAdded:      0,
			LinesSubtracted: 2,
			Name:            "PetersWalker",
			Date:            firstDate,
			CommitHash:      "2528f600f73947495c7396a0d6d5ff2f1a4d343c",
		},
		{
			Filepath:        "",
			LinesAdded:      0,
			LinesSubtracted: 0,
			Name:            "",
			Date:            tomorrow,
			CommitHash:      "",
		},
		{
			Filepath:        "cli/wrapper.go",
			LinesAdded:      0,
			LinesSubtracted: 1,
			Name:            "PetersWalker",
			Date:            twoFromNow,
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
