package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ConfigItem struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Key         string   `json:"key"`
	Cmd         []string `json:"cmd"`
}
type Config struct {
	ConfigItems []ConfigItem `json:"items"`
	Heading     string       `json:"heading"`
}

func GetConfig() Config {
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload Config
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return payload
}
