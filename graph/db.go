package graph

import (
	"os"

	"github.com/mateors/mcb"
)

// var db *mcb.DB
func Connect() *mcb.DB {

	host := os.Getenv("HOST")
	username := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")
	//bucket := os.Getenv("BUCKET")
	//scope := os.Getenv("SCOPE")

	db := mcb.Connect(host, username, password, false)
	res, err := db.Ping()
	if err != nil {
		panic(res)
	}
	return db
}
