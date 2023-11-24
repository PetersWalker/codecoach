package commits

import "time"

type Stats struct {
	Filepath        string
	LinesAdded      int
	LinesSubtracted int
	Name            string
	Date            time.Time
	CommitHash      string
}

type RawStats struct {
	Filepath        string
	LinesAdded      string
	LinesSubtracted string
	Name            string
	Date            time.Time
	CommitHash      string
}

type RawFile struct {
	FilePath   string
	Added      string
	Subtracted string
}

type RawCommit struct {
	CommitHash string
	Author     string
	Date       string
	Files      []RawFile
}

type LogOptions struct {
	AllLogs bool
}
