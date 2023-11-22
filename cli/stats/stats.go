package stats

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
