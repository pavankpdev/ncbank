package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pavankpdev/ncbank/api"
	db "github.com/pavankpdev/ncbank/db/sqlc"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://pavan:pavan@localhost:5432/ncbank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to the DB", err)
	}
	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot Start server", err)
	}
}
