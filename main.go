package main

import (
	"fmt"
	"log"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"

	"github.com/the-gigi/multi-git/cmd"
)

const (
	version = "v0.9.4"
)

var (
	gitTag         = ""
	buildTimestamp = ""
)

func main() {
	if gitTag != "" {
		fmt.Println("Git tag:", gitTag)
	}

	if buildTimestamp != "" {
		fmt.Println("Built at:", buildTimestamp)
	}

	fmt.Println("Current version is: ", version)

	v := semver.MustParse(version[1:])
	latest, err := selfupdate.UpdateSelf(v, "the-gigi/multi-git")
	if err != nil {
		log.Fatalf("Binary update failed: %v", err)
		return
	} else {
		if latest.Version.String() != version {
			fmt.Println("Updated version to:", latest.Version)
		}
	}

	cmd.Execute()
}
