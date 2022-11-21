package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSourse = "postgresql://root:devsecret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	connURI := dbSourse
	if os.Getenv("POSTGRES_URI") != "" {
		connURI = os.Getenv("POSTGRES_URI")
	}

	conn, err := sql.Open(dbDriver, connURI)
	if err != nil {
		log.Fatal("DB connect", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
