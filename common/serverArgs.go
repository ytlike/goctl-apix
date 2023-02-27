package common

import (
	"fmt"
	"qbq-open-platform/common/compile"
)

func ExecArgs(args []string) bool {
	if len(args) >= 2 {
		arg := args[len(args)-1]
		if arg == "version" || arg == "v" {
			v := GetVersion()
			fmt.Printf("Version: %s\nGitBranch: %s\nGitTag：%s\nBuild Date: %s\nGo Version: %s\nOS/Arch: %s\n", v.Version, v.GitBranch, v.GitTag, v.BuildDate, v.GoVersion, v.Platform)
		} else if arg == "debug" {
			pId, err := GetPid(fmt.Sprintf("/opt/app/%s", compile.ServerName))
			if err != nil {
				panic(err.Error())
			} else {
				_, err = RunCommand(fmt.Sprintf("dlv attach %v --headless=true --api-version=2 --accept-multiclient --listen=:5005", pId))
				if err != nil {
					panic(err.Error())
				}
			}
		} else if arg == "start" {
			return true
		}
	}
	return false
}
