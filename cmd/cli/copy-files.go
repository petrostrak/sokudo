package main

import (
	"embed"
	"io/ioutil"
)

var (
	//go:embed templates
	templateFS embed.FS
)

func copyFileFromTemplate(templatePath, targetFile string) error {
	// check to ensure file does not already exist

	data, err := templateFS.ReadFile(templatePath)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile(data, targetFile)
	if err != nil {
		exitGracefully(err)
	}

	return nil
}

func copyDataToFile(data []byte, to string) error {
	err := ioutil.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
