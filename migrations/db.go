package migrations

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func CreateDB() {
	db, err := sql.Open("sqlite3", "./ReefDB.db")
	if err != nil {
		log.Panic(err)
	}
	CreateTableParameter(db)
	CreateTableTanks(db)
	InsertParameter(db)

}

func CreateTableParameter(db *sql.DB) {
	stmt := "CREATE TABLE IF NOT EXISTS parameter (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL," +
		"name STRING," +
		"unit VARCHAR(7)," +
		"formula VARCHAR(7)," +
		"optimal_min DOUBLE," +
		"optimal_max DOUBLE," +
		"min DOUBLE," +
		"max DOUBLE," +
		"created DATETIME DEFAULT CURRENT_TIMESTAMP," +
		"updated DATETIME DEFAULT CURRENT_TIMESTAMP," +
		"enabled INTEGER DEFAULT 1)"
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

func InsertParameter(db *sql.DB) {
	sql := "INSERT INTO parameter (name,unit,formula,optimal_min,optimal_max,min,max) VALUES (?, ?, ?, ?, ?, ?, ?)"
	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()
	stmt.Exec("Nitrate", "mg/l", "NO3", 2, 10, 0, 20)
	stmt.Exec("Phosphate", "mg/l", "PO4", 0.01, 0.1, 0, 0.2)
	stmt.Exec("Calcium", "mg/l", "Ca", 400, 450, 360, 480)
	stmt.Exec("Magnesium", "mg/l", "Mg", 1280, 1350, 1100, 1400)
	stmt.Exec("Alkalinity", "°dH", "", 6.8, 8.5, 5, 20)
	stmt.Exec("Salinity", "g/cm3", "", 1.0233, 1.0233, 1.021, 1.024)
	stmt.Exec("Temperature", "°dH", "", 24, 26, 23.5, 28.3)
	tx.Commit()
}

func CreateTableTanks(db *sql.DB) {
	stmt := "CREATE TABLE IF NOT EXISTS tanks (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL," +
		"name STRING," +
		"volume INTEGER," +
		"location INTEGER," +
		"description TEXT," +
		"manufacturer STRING," +
		"model STRING," +
		"started DATETIME," +
		"created DATETIME DEFAULT CURRENT_TIMESTAMP," +
		"updated DATETIME DEFAULT CURRENT_TIMESTAMP)"
	_, err := db.Exec(stmt)
	if err != nil {
		log.Panic(err)
	}
	stmt = "CREATE UNIQUE INDEX idx_tanks_id ON tanks(id);"
	_, err = db.Exec(stmt)
	if err != nil {
		log.Panic(err)
	}
}
