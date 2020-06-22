package db

import (
	"log"
	"os"
	"github.com/go-pg/pg/v9"

	controller "github.com/amirsorouri00/arvandiscount/controller"
)


// Connecting to DB OK
func ConnectDB() *pg.DB {
	opts := &pg.Options {
		User: "go_db",
		Password: "123123",
		Addr: "localhost:5432",
		Database: "go_db",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	log.Printf("Connected to DB")
	
	// Pass DB Connection to the controller
	controller.InitiateDB(db)
	controller.InitiateSeed()
	controller.CreateTables(db)

	return db
}


