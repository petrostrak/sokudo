package sokudo

import (
	"log"

	"github.com/gobuffalo/pop"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (s *Sokudo) popConnect() (*pop.Connection, error) {
	tx, err := pop.Connect("development")
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *Sokudo) CreatePopMigration(up, down []byte, migrationName, migrationType string) error {
	var migrationPath = s.RootPath + "/migrations"
	err := pop.MigrationCreate(migrationPath, migrationName, migrationType, up, down)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sokudo) PopMigrateUp(tx *pop.Connection) error {
	var migrationPath = s.RootPath + "/migrations"

	fm, err := pop.NewFileMigrator(migrationPath, tx)
	if err != nil {
		return err
	}

	if err = fm.Up(); err != nil {
		return err
	}

	return nil
}

func (s *Sokudo) PopMigrateDown(tx *pop.Connection, steps ...int) error {
	var migrationPath = s.RootPath + "/migrations"

	step := 1
	if len(steps) > 0 {
		step = steps[0]
	}

	fm, err := pop.NewFileMigrator(migrationPath, tx)
	if err != nil {
		return err
	}

	if err = fm.Down(step); err != nil {
		return err
	}

	return nil
}

func (s *Sokudo) PopMigrateReset(tx *pop.Connection) error {
	var migrationPath = s.RootPath + "/migrations"

	fm, err := pop.NewFileMigrator(migrationPath, tx)
	if err != nil {
		return err
	}

	if err = fm.Reset(); err != nil {
		return err
	}

	return nil
}

func (s *Sokudo) MigrateUp(dsn string) error {
	m, err := migrate.New("file://"+s.RootPath+"/migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil {
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

	if err := m.Down(); err != nil {
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

	if err := m.Steps(n); err != nil {
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

	if err := m.Force(-1); err != nil {
		return err
	}

	return nil
}
