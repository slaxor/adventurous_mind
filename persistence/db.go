package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func New() *sql.DB {
	var dbuser = os.ExpandEnv("$DBUSER")
	var dbhost = os.ExpandEnv("$DBHOST")
	var dbname = os.ExpandEnv("$DBNAME")
	var db, err = sql.Open("postgres", fmt.Sprintf("postgres://%s@%s/%s?sslmode=disable", dbuser, dbhost, dbname))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
