package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func migrateDb() {

	// Get a database handle.
	driver, err := mysql.WithInstance(DB, &mysql.Config{})
	if err != nil {
		log.Fatal("Could not create migration driver", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql", driver)
	if err != nil {
		log.Fatal("Could not start migration", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Could not apply migration", err)
	}
}
