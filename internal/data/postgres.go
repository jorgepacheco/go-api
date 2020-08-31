package data

import (
	"database/sql"
	"os"

	// registering database driver
	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {

	// postgres://gopher:gopher@127.0.0.1:5432/microblog?sslmode=disable"
	uri := os.Getenv("DATABASE_URI")

	if uri == "" {
		uri = "postgres://admin:admin@127.0.0.1:5432/accounts?sslmode=disable"
	}
	return sql.Open("postgres", uri)
}
