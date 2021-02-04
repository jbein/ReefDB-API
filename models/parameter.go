package models

import (
	_ "github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Parameter struct {
	Id         uint      `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Unit       string    `json:"unit" db:"unit"`
	Formula    string    `json:"formula" db:"formula"`
	OptimalMin float32   `json:"optimal_min" db:"optimal_min"`
	OptimalMax float32   `json:"optimal_max" db:"optimal_max"`
	Min        float32   `json:"min" db:"min"`
	Max        float32   `json:"max" db:"max"`
	Created    time.Time `json:"created" db:"created"`
	Updated    time.Time `json:"updated" db:"updated"`
	Enabled    bool      `json:"enabled" db:"enabled"`
}
