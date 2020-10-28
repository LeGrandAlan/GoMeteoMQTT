package main

import (
	"./pubsub/subscribers"
	"fmt"
	"time"
)

func main() {

	pool := subscribers.RedisConnect()

	/*captorValues1, _ := subscribers.ScanByAirportAndType(pool, "BDX", "wind")
	fmt.Println(captorValues1)*/

	// => du 25 matin 00:01 au 25 soir 23:59 pour l'aéroport de bordeaux capteur de type vent
	start := time.Date(2020, 10, 27, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 10, 30, 0, 0, 0, 0, time.UTC)
	captorValues2 := subscribers.ScanByAirportAndTypeAndDate(pool, "BDX", "wind", start, end)
	fmt.Println(captorValues2)

	/*captorValues3, _ := subscribers.HGetAll(pool, "goMeteoMQTT:captorValues:BDX:wind325697600936214817")
	fmt.Println(captorValues3)*/

}
