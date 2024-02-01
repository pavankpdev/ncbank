package db

import (
	"database/sql"
	"github.com/pavankpdev/ncbank/util"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://pavan:pavan@localhost:5432/ncbank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to the DB", err)
	}

	testQueries = New(testDB)

	m.Run()
}
