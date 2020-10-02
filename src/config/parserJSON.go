package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func ConfigFileToMap(path string) map[string]interface{} {
	var data interface{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("error:", err)
	}
	return data.(map[string]interface{})
}

func ConfigFileToArray(path string) []interface{} {
	var data interface{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("error:", err)
	}
	return data.([]interface{})
}
