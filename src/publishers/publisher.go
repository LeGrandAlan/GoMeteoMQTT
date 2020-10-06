package publishers

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"strconv"
	"time"
)

func randomValue(min, max float64) float64 {

	return min + rand.Float64()*(max-min)

}

func PublishValue(client mqtt.Client, airportId, captorType string, captorId int, min, max float64) {

	value := randomValue(min, max)
	currentTime := time.Now()
	msg := currentTime.Format("2006-01-02 03:04:05") + ";" +
		airportId + ";" +
		captorType + ";" +
		strconv.Itoa(captorId) + ";" +
		fmt.Sprintf("%.2f", value)
	topic := "/" + airportId + "/" + captorType
	client.Publish(topic, 2, false, msg)

}
