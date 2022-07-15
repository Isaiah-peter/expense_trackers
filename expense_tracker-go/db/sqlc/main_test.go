package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Isaiah-peter/expense_tracker/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("unable to load env", err)
	}
	conn, err := sql.Open(config.DBdriver, config.DBsource)
	if err != nil {
		log.Fatal("can not connect to the database: ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
