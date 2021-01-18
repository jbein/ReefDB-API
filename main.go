package main

import (
	"ReefDB/routes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func main() {
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
