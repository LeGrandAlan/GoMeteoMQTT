package publishers

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
)

func randomValue(min, max float64) float64 {

	return min + rand.Float64()*(max-min)

}

func PublishValue(client mqtt.Client, min, max float64, topic string) {

	value := randomValue(min, max)
	msg := fmt.Sprintf("%.2f", value)
	client.Publish(topic, 2, false, msg)

}
