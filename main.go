package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sharvan/simplebank/api"
	db "github.com/sharvan/simplebank/db/sqlc"
	"github.com/sharvan/simplebank/utils"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalln("Connot load the config")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("Connot coneect to database")
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalln("Error Initializing Errror")
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalln("Connot Start server")
	}
}
