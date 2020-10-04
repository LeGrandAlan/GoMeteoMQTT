package main

import (
	"../src/config"
	"../src/publishers"
	"../src/utils"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jasonlvhit/gocron"
	"time"
)

func main() {

	conf := config.ConfigFileToMap("./config/config.json")
	res := config.ConfigFileToArray("./config/publisher.json")

	for _, object := range res {
		publisher := publishers.MakeFromMap(object.(map[string]interface{}))
		uri := fmt.Sprintf("%v", conf["Protocol"]) + "://" +
			fmt.Sprintf("%v", conf["Host"]) + ":" +
			fmt.Sprintf("%v", conf["Port"])

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
