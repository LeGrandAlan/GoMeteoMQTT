package utils

import (
	"../configUtils"
	"fmt"
)

func GetURIFromConf() string {

	conf := configUtils.ConfigFileToMap("./configUtils/configUtils.json")

	uri := fmt.Sprintf("%v", conf["Protocol"]) + "://" +
		fmt.Sprintf("%v", conf["Host"]) + ":" +
		fmt.Sprintf("%v", conf["Port"])

	return uri

}
