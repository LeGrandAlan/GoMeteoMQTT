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

		go executeCronJob(publishers.PublishValue, publisherClient, publisher.Min, publisher.Max)
	}

	for {
		time.Sleep(100 * time.Millisecond)
	}

}

func executeCronJob(task func(client mqtt.Client, min, max float64), client mqtt.Client, min, max float64) {

	_ = gocron.Every(1).Second().Do(task, client, min, max)
	<-gocron.Start()

}
