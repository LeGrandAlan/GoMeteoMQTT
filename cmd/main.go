package main

import (
	"../src/config"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello world")
	id, _ := uuid.NewRandom()
	fmt.Println(id)
	res := config.ReadConfigFile("./config/config.json").(map[string]interface{})
	fmt.Println(res["Host"])
}
