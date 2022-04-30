package main

import (
	"database/sql"
	"log"

	"github.com/imonasterio/go-sql-docker-aws/api"
	db "github.com/imonasterio/go-sql-docker-aws/db/sqlc"
	"github.com/imonasterio/go-sql-docker-aws/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot conect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
