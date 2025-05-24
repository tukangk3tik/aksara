package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
const DB_DRIVER = "postgres"
const DB_SOURCE_TEST = "postgres://postgres:aksara2025@localhost:5432/aksara_test?sslmode=disable"

func TestMain(m *testing.M) {
	conn, err := sql.Open(DB_DRIVER, DB_SOURCE_TEST)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
