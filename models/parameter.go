package models

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
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

func GetParameter(c *gin.Context) {

}

func GetAllParameter(c *gin.Context) {
	db, err := sql.Open("sqlite3", "./ReefDB.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	sql := "SELECT * FROM parameter"
	rows, err := db.Query(sql)
	fmt.Println(rows)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rows)
}
