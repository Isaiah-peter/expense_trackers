package main

import (
	"database/sql"
	"log"

	"github.com/Isaiah-peter/expense_tracker/api"
	db "github.com/Isaiah-peter/expense_tracker/db/sqlc"
	"github.com/Isaiah-peter/expense_tracker/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBdriver, config.DBsource)
	if err != nil {
		log.Fatal("can not connect to the database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.Serverddress)
	if err != nil {
		log.Fatal("fails to start server", err)
	}
}
