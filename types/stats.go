package types

import "time"

type Stats struct {
	Filepath        string
	LinesAdded      string
	LinesSubtracted string
	Name            string
	Date            time.Time
	Commit          string
}
