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
		log.Println("error runing migration up:", err)
		return err
	}

	return nil
}

func (s *Sokudo) MigrateDownAll(dsn string) error {
	m, err := migrate.New("file://"+s.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Down(); err != nil {
		log.Println("error runing migration down:", err)
		return err
	}

	return nil
}

func (s *Sokudo) Steps(n int, dsn string) error {
	m, err := migrate.New("file://"+s.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Steps(n); err != nil {
		return err
	}

	return nil
}

func (s *Sokudo) MigrateForce(dsn string) error {
	m, err := migrate.New("file://"+s.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Force(-1); err != nil {
		return err
	}

	return nil
}
