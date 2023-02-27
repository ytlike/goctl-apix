package common

import (
	"fmt"
	"runtime"
)

var (
	version   string
	gitBranch string
	gitTag    string
	buildDate string
)

type Version struct {
	Version   string `json:"version"`
	GitBranch string `json:"gitBranch"`
	GitTag    string `json:"gitTag"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
	Compiler  string `json:"compiler"`
	Platform  string `json:"platform"`
}

func GetVersion() *Version {
	return &Version{
		Version:   version,
		GitBranch: gitBranch,
		GitTag:    gitTag,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
