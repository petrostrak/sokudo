package sokudo

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
)

func (s *Sokudo) MigrateUp(dsn string) error {
	m, err := migrate.New("file://"+s.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Up(); err != nil {
		log.Println("error runing migration:", err)
		return err
	}

	return nil
}
