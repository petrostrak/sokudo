package sokudo

type initPaths struct {
	rootPath    string
	folderNames []string
}

type cookieConfig struct {
	name     string
	lifetime string
	persists string
	secure   string
	domain   string
}
