package main

import (
	"../src/config"
	"../src/publishers"
	"fmt"
	"time"
)

/*func main() {
	fmt.Println("Hello world")
	id, _ := uuid.NewRandom()
	fmt.Println(id)
	res := config.ConfigFileToMap("./config/config.json")
	fmt.Println(res["Host"])
}*/

func main() {
	fmt.Println("TEMPERATURE CAPTOR")
	publishers.Client = publishers.Connect("tcp://localhost:1883", "my-client-id")

	res := config.ConfigFileToArray("./config/publisher.json")
	for _, object := range res {
		publisher := publishers.MakeFromMap(object.(map[string]interface{}))
		fmt.Println(publisher)
	}

	publishers.DoEvery(1000 * time.Millisecond, publishers.PublishValue)
}
