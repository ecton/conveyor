package main

import (
	"database/sql"
	"fmt"

	_ "github.com/nilobject/conveyor/Godeps/_workspace/src/github.com/lib/pq"
	_ "github.com/nilobject/conveyor/Godeps/_workspace/src/github.com/mattes/migrate/driver/postgres"
	"github.com/nilobject/conveyor/Godeps/_workspace/src/github.com/mattes/migrate/migrate"
)

func dbURI() string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", config().pg.username, config().pg.password, config().pg.host, config().pg.port, config().pg.databaseName, config().pg.sslMode)
}

func db() (*sql.DB, error) {
	return sql.Open("postgres", dbURI())
}

func migrateDatabaseUp() error {
	logger().Info("Connection URL: %v", dbURI())
	allErrors, ok := migrate.UpSync(dbURI(), "./migrations")
	if !ok {
		return fmt.Errorf("Error performing migrations: %v", allErrors)
	}
	return nil
}

func migrateDatabaseDown() error {
	allErrors, ok := migrate.DownSync(dbURI(), "./migrations")
	if !ok {
		return fmt.Errorf("Error performing migrations: %v", allErrors)
	}
	return nil
}
