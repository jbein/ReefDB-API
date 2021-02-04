package main

import (
	"ReefDB-API/controllers"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("sqlite3", "./ReefDB.db")
	if err != nil {
		fmt.Printf("DB: %v\n", err)
	}

	//migrations.CreateDB()

	router := gin.Default()
	tankGroup := router.Group("tank")
	{
		th := controllers.NewTankHandler(db)
		tankGroup.GET(":id", th.GetTank)
		tankGroup.GET("", th.GetTankAll)
		tankGroup.POST("", th.PostTank)
	}
	parameterGroup := router.Group("parameter")
	{
		ph := controllers.NewParameterHandler(db)
		parameterGroup.GET("", ph.GetParameterAll)
		parameterGroup.GET(":id", ph.GetParameter)
		parameterGroup.POST("", ph.PostParameter)

	}
	_ = router.Run(":3000")

}

func WriteToIFX(p *write.Point) (err error) {
	client := influxdb2.NewClient("https://maui.janbein.de:8086", "ReefDB-API:Ph0sphat3!")
	writeAPI := client.WriteAPIBlocking("", "ReefDB-API")
	err = writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Printf("Write error: %s\n", err.Error())
	}
	return err
}
