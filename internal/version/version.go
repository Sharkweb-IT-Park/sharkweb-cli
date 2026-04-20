package version

import "fmt"

var (
	Version   = "dev"
	Commit    = "none"
	BuildTime = "unknown"
)

func Short() string {
	return Version
}

func Info() string {
	return fmt.Sprintf(
		"Sharkweb CLI\nVersion: %s\nCommit: %s\nBuildTime: %s",
		Version,
		Commit,
		BuildTime,
	)
}
