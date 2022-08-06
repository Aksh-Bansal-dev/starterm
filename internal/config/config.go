package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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
	home, _ := os.UserHomeDir()
	content, err := ioutil.ReadFile(home + "/.starterm.json")
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
