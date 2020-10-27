package main

import (
	"./models"
	"./pubsub/configUtils"
	"./pubsub/subscribers"
)

func main() {

	pool := subscribers.RedisConnect()

	datas := configUtils.ConfigFileToArray("./cmd/pubsub/config/airports.json")

	for _, object := range datas {
		airport := models.AirportMapper(object.(map[string]interface{}))

		_ = subscribers.HSetAirportIfNoExists(pool, airport)
	}

	_, _ = subscribers.ScanAirports(pool)

}
