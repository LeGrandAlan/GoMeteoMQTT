package main

import (
	"./pubsub/configUtils"
	"./pubsub/utils"
	"./pubsub/subscribers"
	"./models"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"time"
)

func main() {
	airports := configUtils.ConfigFileToArray("./cmd/examples/airports.json")

	topics := GetTopicAirport(airports)

	subscriberId := utils.GetEnvOrElse("subscriber_id", 3000)

	client := utils.Connect(utils.GetURIFromConf(), subscriberId)

	client.SubscribeMultiple(topics, func(client mqtt.Client, message mqtt.Message) {
		content := string(message.Payload())
		date := strings.Split(strings.Split(content, ";")[0], " ")[0]
		airportCity := strings.Split(message.Topic(), "/")[1]
		from := strings.Split(message.Topic(), "/")[2]
		filename := airportCity + "-" + date + "-" + from + ".csv"

		go fmt.Println(fmt.Sprintf("{ date: %s, city: %s, from: %s, content: %s}", date, airportCity, from, content))

		go subscribers.PrepareFile(filename).Write(content)
	})

	for {
		time.Sleep(100 * time.Millisecond)
	}
}

func GetTopicAirport(airports []interface{}) map[string]byte {
	topics := make(map[string]byte)

	for _, airport := range airports {
		topics["/"+models.AirportMapper(airport.(map[string]interface{})).Name+"/+"] = 0
	}

	return topics
}
