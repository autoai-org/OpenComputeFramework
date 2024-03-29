package main

import (
	"ocf/bin/cmd"
	"ocf/internal/common"
)

var (
	// Populated during build
	version     = "dev"
	commitHash  = "?"
	buildDate   = ""
	buildSecret = ""
)

func main() {
	common.JSONVersion.Version = version
	common.JSONVersion.Commit = commitHash
	common.JSONVersion.Date = buildDate
	cmd.Execute()
}
