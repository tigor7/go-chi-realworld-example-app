package main

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalln("Failed to connect to DB:", err.Error())
	}
	if err := runMigrations(db.DB); err != nil {
		log.Fatalln("Failed to migrate SQL files:", err.Error())
	}

	fmt.Println("hello world")
}

func connectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "postgresql://localhost:password@127.0.0.1:5432/realworld?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func runMigrations(instance *sql.DB) error {
	driver, err := postgres.WithInstance(instance, &postgres.Config{})
	if err != nil {
		return err
	}
	rawPath, err := filepath.Abs("migrations")
	if err != nil {
		return err
	}
	path := "file:///" + filepath.ToSlash(rawPath)
	m, err := migrate.NewWithDatabaseInstance(
		path,
		"postgres", driver)
	if err != nil {
		return err
	}
	return m.Up()
}
