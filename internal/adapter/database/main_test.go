package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/fbriansyah/bank-ina-test/util"
	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
	testDB      *sql.DB
	testAdapter DatabaseAdapter
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	testAdapter = NewDatabaseAdapter(testDB)

	os.Exit(m.Run())
}
