package sokudo

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/petrostrak/sokudo/render"
)

const (
	version = "1.0.0"
)

type Sokudo struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	Render   *render.Render
	config
}

type config struct {
	port     string
	rendeder string
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

	// create loggers
	infoLog, errorLog := s.startLoggers()
	s.InfoLog = infoLog
	s.ErrorLog = errorLog
	s.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	s.Version = version
	s.RootPath = rootPath
	s.Routes = s.routes().(*chi.Mux)

	s.config = config{
		port:     os.Getenv("PORT"),
		rendeder: os.Getenv("RENDERER"),
	}

	s.Render = s.createRenderer(s)

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

// ListenAndServe starts the web server
func (s *Sokudo) ListenAndServe() {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     s.ErrorLog,
		Handler:      s.routes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	s.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))
	if err := srv.ListenAndServe(); err != nil {
		s.ErrorLog.Fatal(err)
	}
}

func (s *Sokudo) checkDotEnv(path string) error {
	err := s.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}

	return nil
}

func (s *Sokudo) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog, errorLog *log.Logger

	infoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func (s *Sokudo) createRenderer(skd *Sokudo) *render.Render {
	myRender := render.Render{
		Renderer: skd.config.rendeder,
		RootPath: skd.RootPath,
		Port:     skd.config.port,
	}

	return &myRender
}
