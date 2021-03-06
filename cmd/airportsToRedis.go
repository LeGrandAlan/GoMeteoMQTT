package main

import (
	"./models"
	"./pubsub/subscribers"
	"./pubsub/utils"
	"fmt"
)

func main() {

	pool := subscribers.RedisConnect()

	datas := utils.ConfigFileToArray("./cmd/pubsub/config/airports.json")

	for _, object := range datas {
		airport := models.AirportMapper(object.(map[string]interface{}))

		_ = subscribers.HSetAirportIfNoExists(pool, airport)
	}

	airports, _ := subscribers.ScanAirports(pool)

	fmt.Println(airports)

}
