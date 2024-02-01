package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pavankpdev/ncbank/api"
	db "github.com/pavankpdev/ncbank/db/sqlc"
	"github.com/pavankpdev/ncbank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to the DB", err)
	}
	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start server", err)
	}
}
