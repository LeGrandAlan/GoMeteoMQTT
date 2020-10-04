package utils

import (
	"../config"
	"fmt"
)

func GetURIFromConf() string {

	conf := config.ConfigFileToMap("./config/config.json")

	uri := fmt.Sprintf("%v", conf["Protocol"]) + "://" +
		fmt.Sprintf("%v", conf["Host"]) + ":" +
		fmt.Sprintf("%v", conf["Port"])

	return uri

}
