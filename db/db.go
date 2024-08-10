package db

import (
	"database/sql"
	"go-event-api/utils"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	if utils.Env.IsDevMode() {
		DB, err = sql.Open("sqlite3", utils.Env.DatabaseURL)
	} else {
		cfg := mysql.Config{
			User:                 utils.Env.DatabaseUser,
			Passwd:               utils.Env.DatabasePassword,
			Net:                  "tcp",
			Addr:                 utils.Env.DatabaseURL,
			ParseTime:            true,
			DBName:               "events-db",
			AllowNativePasswords: true,
		}
		// Get a database handle.
		DB, err = sql.Open("mysql", cfg.FormatDSN())
	}

	if err != nil {
		log.Panic("Could not connect to the DB ", err)

	}

	log.Print("DB connection was established")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	if !utils.Env.IsDevMode() {
		_, err := DB.Exec(`SHOW DATABASES LIKE 'events-db';`)
		if err != nil {
			log.Panic("HELLO", err)
		}
	}

	createTables()
}

func createTables() {
	createTableUser := `
	CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
	);
	`

	_, err := DB.Exec(createTableUser)

	if err != nil {
		log.Panic("could not create users table", err)
	}

	createTableEvents := `
	CREATE TABLE IF NOT EXISTS events (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    location VARCHAR(255) NOT NULL,
    dateTime DATETIME NOT NULL,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	_, err = DB.Exec(createTableEvents)

	if err != nil {
		log.Panic("could not create events table", err)
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
    id INT PRIMARY KEY AUTO_INCREMENT,
    event_id INT,
    user_id INT,
    FOREIGN KEY (event_id) REFERENCES events(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
	);`
	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		log.Panic("could not create registrations table", err)
	}

}
