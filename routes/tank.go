package routes

import (
	"ReefDB/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TankAdd(c *gin.Context) {
	newTank := models.Tank{}
	err := c.ShouldBindJSON(&newTank)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error binding JSON!")
	}
	_, ok := models.AddNewTank(newTank)
	if !ok {
		c.String(http.StatusInternalServerError, "Tank not created!")
		return
	}
	c.JSON(http.StatusCreated, newTank)
}

func TankShowAll(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetAllTanks())
}

func TankShow(c *gin.Context) {
	id := c.Param("id")
	id2, err := strconv.ParseUint(id, 10, 8)
	if err != nil {
		fmt.Println(err.Error())
	}
	tank, ok := models.GetTank(uint(id2))
	if !ok {
		c.String(http.StatusNotFound, "Tank not found!")
		return
	}
	c.JSON(http.StatusOK, tank)
}
