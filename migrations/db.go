package migrations

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func createDB() {
	db, err := sql.Open("sqlite3", "./ReefDB.db")
	if err != nil {
		log.Panic(err)
	}
	createTableParameter(db)

}

func createTableParameter(db *sql.DB) {
	stmt := "CREATE TABLE IF NOT EXISTS parameter (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL," +
		"name STRING," +
		"unit VARCHAR(7)," +
		"formula VARCHAR(7)," +
		"optimal_min DOUBLE," +
		"optimal_max DOUBLE," +
		"min DOUBLE," +
		"max DOUBLE)"
	_, err := db.Exec(stmt)
	if err != nil {
		log.Panic(err)
	}
	stmt = "CREATE UNIQUE INDEX idx_parameter_id ON parameter(id);"
	_, err = db.Exec(stmt)
	if err != nil {
		log.Panic(err)
	}
}
