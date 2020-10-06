package utils

import (
	"../configUtils"
	"fmt"
)

func GetURIFromConf() string {

	conf := configUtils.ConfigFileToMap("./cmd/pubsub/config/config.json")

	uri := fmt.Sprintf("%v", conf["Protocol"]) + "://" +
		fmt.Sprintf("%v", conf["Host"]) + ":" +
		fmt.Sprintf("%v", conf["Port"])

	return uri

}
