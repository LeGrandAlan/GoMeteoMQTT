package main

import (
	"./pubsub/subscribers"
	"fmt"
)

var (
// Pool *redis.Pool
)

func main() {

	Pool = subscribers.RedisConnect()

	captorValues1, _ := subscribers.ScanByAirportAndType(Pool, "BDX", "wind")
	fmt.Println(captorValues1)

	/*captorValues2, _ := subscribers.HGetAll(Pool, "goMeteoMQTT:captorValues:BDX:wind325697600936214817")
	fmt.Println(captorValues2)*/

}
