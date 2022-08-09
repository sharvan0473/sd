package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5435/simplebank?sslmode=disable"
)

var testQueries *Queries

func TestMain(t *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to sql")
	}
	testQueries = New(conn)
	os.Exit(t.Run())
}
