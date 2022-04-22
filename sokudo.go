package sokudo

import (
	"fmt"

	"github.com/joho/godotenv"
)

const (
	version = "1.0.0"
)

type Sokudo struct {
	AppName string
	Debug   bool
	Version string
}

func (s *Sokudo) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := s.Init(pathConfig)
	if err != nil {
		return err
	}

	err = s.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	// read .env
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	return nil
}

func (s *Sokudo) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		// crate folder if it doesn't exist
		err := s.CreateDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Sokudo) checkDotEnv(path string) error {
	err := s.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}

	return nil
}
