package main

import (
	"./pubsub/configUtils"
	"./pubsub/publishers"
	"./pubsub/utils"
	"bufio"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jasonlvhit/gocron"
	"os"
	"time"
)

func main() {

	res := configUtils.ConfigFileToArray("./configUtils/publisher.json")

	for _, object := range res {
		publisher := publishers.MakeFromMap(object.(map[string]interface{}))
		uri := utils.GetURIFromConf()

		publisherClient := utils.Connect(uri, publisher.Id)

		go executeCronJob(publishers.PublishValue, publisherClient, publisher.AirportId, publisher.Type, publisher.Id, publisher.Min, publisher.Max)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Print("\nEnter :q to quit\n")

	for {
		time.Sleep(100 * time.Millisecond)
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		if text == ":q\n" {
			os.Exit(0)
		}
	}

}

func executeCronJob(
	task func(client mqtt.Client, airportId, captorType string, captorId int, min, max float64),
	client mqtt.Client, airportId, captorType string, captorId int, min, max float64) {

	_ = gocron.Every(10).Second().Do(task, client, airportId, captorType, captorId, min, max)
	<-gocron.Start()

}
