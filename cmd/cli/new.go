package main

import (
	"log"
	"strings"
)

func doNew(appName string) {
	appName = strings.ToLower(appName)

	// sanitize the application name (convert url to single word)
	if strings.Contains(appName, "/") {
		exploded := strings.SplitAfter(appName, "/")
		appName = exploded[len(exploded)-1]
	}

	log.Println("App name is", appName)

	// git clone the skeleton application

	// remove the .git directory

	// create a ready to go .env file

	// create a Makefile

	// update the go.mod file

	// update existing .go files with correct name/imports

	// run go mod tidy in the project directory
}
