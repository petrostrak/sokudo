package main

import (
	"errors"
	"os"

	"github.com/fatih/color"
	"github.com/petrostrak/sokudo"
)

const (
	version = "1.0.0"
)

var (
	skd sokudo.Sokudo
)

func main() {
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully(err)
	}
}

func validateInput() (string, string, string, error) {
	var arg1, arg2, arg3 string
	if len(os.Args) > 1 {
		arg1 = os.Args[1]

		if len(os.Args) >= 3 {
			arg2 = os.Args[2]
		}

		if len(os.Args) >= 4 {
			arg3 = os.Args[3]
		}
	} else {
		color.Red("error: command required")
		showHelp()
		return "", "", "", errors.New("command required")
	}

	return arg1, arg2, arg3, nil
}

func showHelp() {
	color.Yellow(`Available commmands:
	help		- show the help commands
	version		- print application version
	`)
}
