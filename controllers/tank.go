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

type TankHandler struct {
	db     *sqlx.DB
	logger *log.Logger
}

func NewTankHandler(db *sqlx.DB) *TankHandler {
	logger := log.New(os.Stdout, "TankHandler", log.LstdFlags)
	return &TankHandler{
		db:     db,
		logger: logger,
	}
}

type PostTank struct {
	Name         string `json:"name" db:"name"`
	Volume       uint   `json:"volume" db:"volume"`
	Location     string `json:"location" db:"location"`
	Description  string `json:"description" db:"description"`
	Manufacturer string `json:"manufacturer" db:"manufacturer"`
	Model        string `json:"model" db:"model"`
}

func (th TankHandler) PostTank(c *gin.Context) {
	input := PostTank{}
	if err := c.ShouldBindJSON(&input); err != nil {
		th.logger.Println("PostTank: ", err.Error())
		c.String(http.StatusBadRequest, "INVALID JSON")
		return
	}
	result, err := th.db.NamedExec("INSERT INTO tank (name, volume, location, description, manufacturer, model) "+
		"VALUES (:name, :volume, :location, :description, :manufacturer, :model)", &input)
	if err != nil {
		th.logger.Println("PostTank: ", err.Error())
		c.String(http.StatusBadRequest, "INSERT FAILD")
		return
	}
	c.JSON(http.StatusOK, result)
}

func (th TankHandler) GetTankAll(c *gin.Context) {
	tanks := make([]models.Tank, 0)
	err := th.db.Select(&tanks, "SELECT * FROM tanks")
	if err != nil {
		th.logger.Println("Ole")
		c.String(http.StatusInternalServerError, "DB ERROR")
		return
	}
	c.JSON(http.StatusOK, tanks)
}

func (th TankHandler) GetTank(c *gin.Context) {
	id := c.Param("id")
	id2, err := strconv.ParseUint(id, 10, 8)
	if err != nil {
		th.logger.Println("Ole")
		c.String(http.StatusBadRequest, "INVALID ID")
		return
	}
	tank := models.Tank{}
	err = th.db.Select(&tank, "SELECT * FROM tanks WHERE id = ?", id2)
	if err != nil {
		th.logger.Println("Ole")
		c.String(http.StatusInternalServerError, "DB ERROR")
		return
	}
	c.JSON(http.StatusOK, tank)
}
