package main

import (
	"./pubsub/utils"
	"bufio"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	uri := utils.GetURIFromConf()
	subscriberClient := utils.Connect(uri, 234)
	subscriberClient.Subscribe("/+/+", 2, test)

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

func test(client mqtt.Client, msg mqtt.Message) {

	strMessage := string(msg.Payload())
	infos := parseMsg(strMessage)
	infosMap := infosToMap(infos)
	fmt.Println(infosMap)

}

func parseMsg(msg string) []string {

	return strings.Split(msg, ";")

}

func infosToMap(infos []string) map[string]interface{} {
	layout := "2006-01-02 03:04:05"

	msgInfos := make(map[string]interface{})
	msgInfos["datetime"], _ = time.Parse(layout, infos[0])
	msgInfos["airportId"] = infos[1]
	msgInfos["captorType"] = infos[2]
	msgInfos["captorId"], _ = strconv.Atoi(infos[3])
	msgInfos["value"], _ = strconv.ParseFloat(infos[4], 64)

	return msgInfos

}
