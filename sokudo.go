package sokudo

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	"github.com/petrostrak/sokudo/cache"
	"github.com/petrostrak/sokudo/render"
	"github.com/petrostrak/sokudo/session"
)

const (
	version = "1.0.0"
)

var (
	myRedisCache *cache.RedisCache
)

type Sokudo struct {
	AppName       string
	Debug         bool
	Version       string
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	RootPath      string
	Routes        *chi.Mux
	Render        *render.Render
	Session       *scs.SessionManager
	DB            Database
	JetViews      *jet.Set
	EncryptionKey string
	Cache         cache.Cache
	config
}

type config struct {
	port        string
	rendeder    string
	cookie      cookieConfig
	sessionType string
	database    databaseConfig
	redis       redisConfig
}

// New reads the .env file, creates our application config, populates the Sokudo type with settings
// based of .env values, and creates necessary folders and files if they don't exist
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

	// connect to database
	if os.Getenv("DATABASE_TYPE") != "" {
		db, err := s.OpenDB(os.Getenv("DATABASE_TYPE"), s.BuildDSN())
		if err != nil {
			errorLog.Println(err)
			os.Exit(1)
		}

		s.DB = Database{
			DataType: os.Getenv("DATABASE_TYPE"),
			Pool:     db,
		}
	}

	if os.Getenv("CACHE") == "redis" || os.Getenv("SESSION_TYPE") == "redis" {
		myRedisCache = s.createClientRedisCache()
		s.Cache = myRedisCache
	}

	s.InfoLog = infoLog
	s.ErrorLog = errorLog
	s.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	s.Version = version
	s.RootPath = rootPath
	s.Routes = s.routes().(*chi.Mux)

	s.config = config{
		port:     os.Getenv("PORT"),
		rendeder: os.Getenv("RENDERER"),
		cookie: cookieConfig{
			name:     os.Getenv("COOKIE_NAME"),
			lifetime: os.Getenv("COOKIE_LIFETIME"),
			persist:  os.Getenv("COOKIE_PERSISTS"),
			secure:   os.Getenv("COOKIE_SECURE"),
			domain:   os.Getenv("COOKIE_DOMAIN"),
		},
		sessionType: os.Getenv("SESSION_TYPE"),
		database: databaseConfig{
			database: os.Getenv("DATABASE_TYPE"),
			dsn:      s.BuildDSN(),
		},
		redis: redisConfig{
			host:     os.Getenv("REDIS_HOST"),
			password: os.Getenv("REDIS_PASSWORD"),
			prefix:   os.Getenv("REDIS_PREFIX"),
		},
	}

	// create session
	sess := session.Session{
		CookieLifetime: s.cookie.lifetime,
		CookiePersist:  s.config.cookie.persist,
		CookieName:     s.config.cookie.name,
		SessionType:    s.config.sessionType,
		CookieDomain:   s.config.cookie.domain,
	}

	switch s.config.sessionType {
	case "redis":
		sess.RedisPool = myRedisCache.Conn
	case "mysql", "mariadb", "postgres", "postresql":
		sess.DBPool = s.DB.Pool
	}

	s.Session = sess.InitSession()
	s.EncryptionKey = os.Getenv("KEY")

	if s.Debug {
		var views = jet.NewSet(
			jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
			jet.InDevelopmentMode(),
		)

		s.JetViews = views
	} else {
		var views = jet.NewSet(
			jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		)

		s.JetViews = views
	}

	s.createRenderer()

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
		Handler:      s.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	defer s.DB.Pool.Close()

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

func (s *Sokudo) createRenderer() {
	myRender := render.Render{
		Renderer: s.config.rendeder,
		RootPath: s.RootPath,
		Port:     s.config.port,
		JetViews: s.JetViews,
		Session:  s.Session,
	}

	s.Render = &myRender
}

func (s *Sokudo) BuildDSN() string {
	var dsn string

	switch os.Getenv("DATABASE_TYPE") {
	case "postgres", "postgresql":
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_NAME"),
			os.Getenv("DATABASE_SSL_MODE"))

		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("%s password=%s", dsn, os.Getenv("DATABASE_PASS"))
		}
	default:
	}

	return dsn
}

func (s *Sokudo) createRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   10000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				s.config.redis.host,
				redis.DialPassword(s.config.redis.password))
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

func (s *Sokudo) createClientRedisCache() *cache.RedisCache {
	return &cache.RedisCache{
		Conn:   s.createRedisPool(),
		Prefix: s.config.redis.prefix,
	}
}
