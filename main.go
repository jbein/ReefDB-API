package main

import (
	"context"
	"database/sql"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./ReefDB.db")
	if err != nil {
		log.Panic(err)
	}
	stmt := "CREATE TABLE IF NOT EXISTS parameter (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL," +
		"name STRING," +
		"unit VARCHAR(7)," +
		"formula VARCHAR(7)," +
		"optimal_min DOUBLE," +
		"optimal_max DOUBLE," +
		"min DOUBLE," +
		"max DOUBLE)"
	_, err = db.Exec(stmt)
	if err != nil {
		log.Panic(err)
	}
	stmt = "CREATE UNIQUE INDEX idx_parameter_id ON parameter(id);"
	_, err = db.Exec(stmt)
	if err != nil {
		log.Panic(err)
	}

	/*
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
	*/
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
