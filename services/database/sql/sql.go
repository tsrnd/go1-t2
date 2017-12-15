package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connect func
func Connect(dlct, user, pass, host, port, name, sslmode string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password='%s' host=%s port=%s dbname=%s sslmode=%s",
		user, pass, host, port, name, sslmode)
	return sql.Open(dlct, connStr)
}
