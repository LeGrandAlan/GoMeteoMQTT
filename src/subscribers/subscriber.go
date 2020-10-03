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
	client := utils.Connect("tcp://localhost:1883", "my-client-id")
	client.Subscribe("/commit", 0, subcribeCallback)
}
