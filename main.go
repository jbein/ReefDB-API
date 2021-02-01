package main

import (
	"ReefDB-API/routes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func main() {
	//migrations.CreateDB()

	router := gin.Default()
	tankGroup := router.Group("tank")
	{
		tankGroup.GET(":id", routes.TankShow)
		tankGroup.POST("", routes.TankAdd)
	}
	tanksGroup := router.Group("tanks")
	{
		tanksGroup.GET("", routes.TankShowAll)
	}
	parameterGroup := router.Group("parameter")
	{
		parameterGroup.GET("", routes.ParameterShowAll)
		parameterGroup.GET(":id", routes.ParameterShow)

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
