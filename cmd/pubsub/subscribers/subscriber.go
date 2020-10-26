package subscribers

import (
	"../utils"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"strings"
)

var (
	Pool *redis.Pool
)

func OnReceive(client mqtt.Client, msg mqtt.Message) {

	strMessage := string(msg.Payload())
	infos := parseMsg(strMessage)
	uniqueId := int(utils.GenUniqueId())
	captorValue := MakeFromArray(infos, uniqueId)

	fmt.Println(uniqueId)

	Pool = RedisConnect()
	_ = Ping(Pool)

	idPrefix := "goMeteoMQTT:captorValues:" + captorValue.AirportId + ":" + captorValue.Type

	_ = HSetCaptorValue(captorValue, idPrefix, strconv.Itoa(uniqueId))

}

func parseMsg(msg string) []string {

	return strings.Split(msg, ";")

}
