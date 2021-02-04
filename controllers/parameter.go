package controllers

import (
	"ReefDB-API/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"strconv"
)

type ParameterHandler struct {
	db     *sqlx.DB
	logger *log.Logger
}

func NewParameterHandler(db *sqlx.DB) *ParameterHandler {
	logger := log.New(os.Stdout, "ParameterHandler", log.LstdFlags)
	return &ParameterHandler{
		db:     db,
		logger: logger,
	}
}

type PostParameter struct {
	Name       string  `json:"name" db:"name"`
	Unit       string  `json:"unit" db:"unit"`
	Formula    string  `json:"formula" db:"formula"`
	OptimalMin float32 `json:"optimal_min" db:"optimal_min"`
	OptimalMax float32 `json:"optimal_max" db:"optimal_max"`
	Min        float32 `json:"min" db:"min"`
	Max        float32 `json:"max" db:"max"`
}

func (ph ParameterHandler) PostParameter(c *gin.Context) {
	input := PostParameter{}
	if err := c.ShouldBindJSON(&input); err != nil {
		ph.logger.Println("PostParameter: ", err.Error())
		c.String(http.StatusBadRequest, "INVALID JSON")
		return
	}
	result, err := ph.db.NamedExec("INSERT INTO parameter (name, unit, formula, optimal_min, optimal_max, min, max) "+
		"VALUES (:name, :unit, :formula, :optimal_min, :optimal_max, :min, :max)", &input)
	if err != nil {
		ph.logger.Println("PostParameter: ", err.Error())
		c.String(http.StatusBadRequest, "INSERT FAILD")
		return
	}
	c.JSON(http.StatusOK, result)
}

func (ph ParameterHandler) GetParameterAll(c *gin.Context) {
	parameters := make([]models.Parameter, 0)
	err := ph.db.Select(&parameters, "SELECT * FROM parameter")
	if err != nil {
		ph.logger.Println("Ole")
		c.String(http.StatusInternalServerError, "DB ERROR")
		return
	}
	c.JSON(http.StatusOK, parameters)
}

func (ph ParameterHandler) GetParameter(c *gin.Context) {
	id := c.Param("id")
	id2, err := strconv.ParseUint(id, 10, 8)
	if err != nil {
		ph.logger.Println("Ole")
		c.String(http.StatusBadRequest, "INVALID ID")
		return
	}
	parameter := models.Parameter{}
	err = ph.db.Select(&parameter, "SELECT * FROM parameter WHERE id = ?", id2)
	if err != nil {
		ph.logger.Println("Ole")
		c.String(http.StatusInternalServerError, "DB ERROR")
		return
	}
	c.JSON(http.StatusOK, parameter)
}
