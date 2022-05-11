package main

import "embed"

//go:embed templates
var (
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
