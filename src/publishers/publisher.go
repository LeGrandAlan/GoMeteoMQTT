package publishers

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"math/rand"
	"time"
)

func randomValue(min, max float64) float64 {

	return min + rand.Float64()*(max-min)

}

func PublishValue(client mqtt.Client, min, max float64) {

	value := randomValue(min, max)
	msg := fmt.Sprintf("%.2f", value)
	client.Publish("test", 2, false, msg)

}

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientId)
	return opts

}

func Connect(brokerURI string, clientId string) mqtt.Client {

	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId + ")...")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected  to broker " + brokerURI + ", with client ID " + clientId + "")
	}
	return client

}
