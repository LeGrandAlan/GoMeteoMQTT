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

		go executeCronJob(publishers.PublishValue, publisherClient, publisher.Min, publisher.Max, publisher.Topic)
	}

	for {
		time.Sleep(100 * time.Millisecond)
	}

}

func executeCronJob(
	task func(client mqtt.Client, min, max float64, topic string),
	client mqtt.Client, min, max float64, topic string) {

	_ = gocron.Every(1).Second().Do(task, client, min, max, topic)
	<-gocron.Start()

}
