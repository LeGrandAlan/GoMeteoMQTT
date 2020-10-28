package main

import (
	"./models"
	"./pubsub/subscribers"
	"./pubsub/utils"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"sync"
	"time"
)

func main() {
	airports := utils.ConfigFileToArray("./cmd/pubsub/config/airports.json")

	topics := GetTopicAirport(airports)

	subscriberId := utils.GetEnvOrElse("subscriber_id", 3000)

	client := utils.Connect(utils.GetURIFromConf(), subscriberId)

	fileMutex := &sync.Mutex{}

	client.SubscribeMultiple(topics, func(client mqtt.Client, message mqtt.Message) {
		content := string(message.Payload())
		date := strings.Split(strings.Split(content, ";")[0], " ")[0]
		airportCity := strings.Split(message.Topic(), "/")[1]
		from := strings.Split(message.Topic(), "/")[2]
		filename := airportCity + "-" + date + "-" + from + ".csv"


		go fmt.Println(fmt.Sprintf("{ date: %s, city: %s, from: %s, content: %s}", date, airportCity, from, content))

		go subscribers.PrepareFile(filename, fileMutex).Write(content)
	})

	for {
		time.Sleep(100 * time.Millisecond)
	}
}

func GetTopicAirport(airports []interface{}) map[string]byte {
	topics := make(map[string]byte)

	for _, airport := range airports {
		topics["/"+models.AirportMapper(airport.(map[string]interface{})).Id+"/+"] = 0
	}

	return topics
}
