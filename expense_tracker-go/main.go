package main

import (
	"database/sql"
	"log"

	"github.com/Isaiah-peter/expense_tracker/api"
	db "github.com/Isaiah-peter/expense_tracker/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dBDriver      = "postgres"
	DBsource      = "postgresql://postgres12:secret@localhost:5342/expense_tracker?sslmode=disable"
	serveraddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dBDriver, DBsource)
	if err != nil {
		log.Fatal("can not connect to the database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serveraddress)
	if err != nil {
		log.Fatal("fails to start server", err)
	}
}
