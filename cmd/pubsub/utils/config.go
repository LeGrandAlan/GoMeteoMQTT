package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetURIFromConf() string {
	conf := ConfigFileToMap("./cmd/pubsub/config/config.json")

	uri := fmt.Sprintf("%v", conf["Protocol"]) + "://" +
		fmt.Sprintf("%v", conf["Host"]) + ":" +
		fmt.Sprintf("%v", conf["Port"])

	return uri
}

func GetEnvOrElse(key string, defaultVal int) int {
	val, present := os.LookupEnv(key)
	if !present {
		return defaultVal
	} else {
		env, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		return env
	}
}
