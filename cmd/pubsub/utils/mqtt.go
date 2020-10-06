package utils

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"strconv"
	"time"
)

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientId)
	opts.SetKeepAlive(3 * time.Second)

	return opts

}

func Connect(brokerURI string, clientId int) mqtt.Client {

	stringClientID := strconv.Itoa(clientId)
	fmt.Println("Trying to connect (" + brokerURI + ", " + stringClientID + ")...")
	opts := createClientOptions(brokerURI, stringClientID)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected  to broker " + brokerURI + ", with client ID " + stringClientID + "")
	}
	return client

}
