package main

import (
	"database/sql"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tigor7/go-chi-realworld-example-app/internal/users"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalln("Failed to connect to DB:", err.Error())
	}
	if err := runMigrations(db.DB); err != nil {
		log.Fatalln("Failed to migrate SQL files:", err.Error())
	}
	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: buildHandler(),
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("Failed to serve the app:", err.Error())
	}
}
func buildHandler() http.Handler {
	r := chi.NewRouter()
	usersHandler := users.NewUsersHandler()
	usersHandler.RegisterRoutes(r)
	return r
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
