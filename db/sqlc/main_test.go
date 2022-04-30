package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/imonasterio/go-sql-docker-aws/util"
	_ "github.com/lib/pq"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:secrets@localhost:5432/easy_bank?sslmode=disable"
// )

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot conect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
