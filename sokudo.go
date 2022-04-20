package sokudo

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
