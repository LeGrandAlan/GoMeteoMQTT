package main

import (
	"../src/config"
	"../src/publishers"
	"../src/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jasonlvhit/gocron"
	"time"
)

func main() {

	res := config.ConfigFileToArray("./config/publisher.json")

	for _, object := range res {
		publisher := publishers.MakeFromMap(object.(map[string]interface{}))
		uri := utils.GetURIFromConf()

		publisherClient := utils.Connect(uri, publisher.Id)

		go executeCronJob(publishers.PublishValue, publisherClient, publisher.AirportId, publisher.Type, publisher.Id, publisher.Min, publisher.Max)
		time.Sleep(100 * time.Millisecond)
	}

	for {
		time.Sleep(100 * time.Millisecond)
	}

}

func executeCronJob(
	task func(client mqtt.Client, airportId, captorType string, captorId int, min, max float64),
	client mqtt.Client, airportId, captorType string, captorId int, min, max float64) {

	_ = gocron.Every(10).Second().Do(task, client, airportId, captorType, captorId, min, max)
	<-gocron.Start()

}
