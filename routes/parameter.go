package routes

import (
	"ReefDB-API/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ParameterShowAll(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetAllParameter())
}

func ParameterShow(c *gin.Context) {
	id := c.Param("id")
	id2, err := strconv.ParseUint(id, 10, 8)
	if err != nil {
		fmt.Println(err.Error())
	}
	parameter, ok := models.GetParameter(int(id2))
	if !ok {
		c.String(http.StatusNotFound, "Parameter not found!")
		return
	}
	c.JSON(http.StatusOK, parameter)
}
