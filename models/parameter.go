package models

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"time"
)

type Parameter struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Unit       string    `json:"unit"`
	Formula    string    `json:"formula"`
	OptimumMin float32   `json:"optimum_min"`
	OptimumMax float32   `json:"optimum_max"`
	Min        float32   `json:"min"`
	Max        float32   `json:"max"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
	Enabled    bool      `json:"enabled"`
}

func AddNewParameter(c *gin.Context) {
	var input Parameter
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func GetParameter(id int) (*Parameter, bool) {
	var parameter []Parameter
	parameter = GetAllParameter()
	return findParameter(parameter, id)
}

func findParameter(parameter []Parameter, id int) (*Parameter, bool) {
	for _, v := range parameter {
		if v.Id == id {
			return &v, true
		}
	}
	return nil, false
}

func GetAllParameter() []Parameter {
	db, err := sql.Open("sqlite3", "./ReefDB.db")
	if err != nil {
		fmt.Printf("DB: %v\n", err)
	}
	sql := "SELECT id,name,unit,formula,optimal_min,optimal_max,min,max,created,updated,enabled FROM parameter"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Printf("Query: %v\n", err)
	}
	var got []Parameter
	for rows.Next() {
		var p Parameter
		rows.Scan(&p.Id, &p.Name, &p.Unit, &p.Formula, &p.OptimumMin, &p.OptimumMax, &p.Min, &p.Max, &p.Created, &p.Updated, &p.Enabled)
		if err != nil {
			fmt.Printf("Scan: %v\n", err)
		}
		got = append(got, p)
	}
	return got
}
