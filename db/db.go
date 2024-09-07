package db

import (
	"database/sql"
	"fmt"
	"go-event-api/utils"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error

	cfg := mysql.Config{
		DBName:               utils.Env.Database,
		Addr:                 utils.Env.DatabaseURL,
		User:                 utils.Env.DatabaseUser,
		Passwd:               utils.Env.DatabasePassword,
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	fmt.Printf("OPEN CONNECTION %v \n", cfg.FormatDSN())
	// Get a database handle.
	DB, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Panic("*** Error: Could not create DB connector. ", err)
		return
	}
	fmt.Println("CONNECTION CREATED")

	err = DB.Ping()

	fmt.Print("TESTING CONNECTION")
	if err != nil {
		log.Panic("*** Error:  Could not connect to DB. ", err)
		return
	}
	log.Print("DB Connection established!")
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	migrateDb()

}
