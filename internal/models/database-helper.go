package models

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDatabase() {
	cfg := mysql.Config{
		User:                 "dac-san-api",
		Passwd:               "nhatnam2002",
		Net:                  "tcp",
		Addr:                 "35.236.191.193:3306",
		DBName:               "dacsandb",
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	// db, err = connectWithConnector()
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

func TaoIdMoi(tenBang string) int {
	var count int
	err := db.QueryRow("SELECT MAX(id) FROM " + tenBang).Scan(&count)
	if err != nil {
		return 1
	}
	return count + 1
}
