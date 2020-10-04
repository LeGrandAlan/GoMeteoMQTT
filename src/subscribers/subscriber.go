package subscribers

import (
	"../utils"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func subcribeCallback(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("Callback")
}

func subscriberMain() {
	uri := utils.GetURIFromConf()
	client := utils.Connect(uri, "my-client-id")
	client.Subscribe("/commit", 0, subcribeCallback)
}
