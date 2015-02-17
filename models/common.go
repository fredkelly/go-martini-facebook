package models

import (
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

var DbMap *gorp.DbMap

// Initialise DB and setup tables
func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@"+os.Getenv("MYSQL_HOST")+"/"+os.Getenv("MYSQL_DBNAME"))
	if err != nil {
		log.Printf("couldn't connect to database: %s", err)
	}

	DbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	DbMap.AddTableWithName(User{}, "users").SetKeys(true, "Id")

	err = DbMap.CreateTablesIfNotExists()
	if err != nil {
		log.Printf("couldn't create tables (%s)", err)
	}

	return DbMap
}

func init() {
	// setup database
	DbMap = initDb()
	// TODO defer DbMap.Db.Close()
}
