package commits

import "time"

type Stats struct {
	Filepath        string    `json:"filePath"`
	LinesAdded      int       `json:"linesAdded"`
	LinesSubtracted int       `json:"linesSubtracted"`
	Name            string    `json:"name"`
	Date            time.Time `json:"date"`
	CommitHash      string    `json:"commitHash"`
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
