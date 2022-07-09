package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dBDriver = "postgres"
	DBsource = "postgresql://postgres12:secret@localhost:5342/expense_tracker?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dBDriver, DBsource)
	if err != nil {
		log.Fatal("can not connect to the database: ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
