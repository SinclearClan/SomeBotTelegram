package config

import (
	"encoding/json"
	"os"
)

// Struct representing the config file for the server, bot and used APIs.
type Config struct {
	Webserver struct {
		Port string `json:"port"`
		Key string `json:"key"`
	} `json:"webserver"`
	Telegram struct {
		Key string `json:"key"`
		Whitelist []int64 `json:"whitelist"`
	} `json:"telegram"`
	Macrodroid struct {
		RestUrl string `json:"rest_url"`
	} `json:"macrodroid"`
	DHL struct {
		Key string `json:"key"`
		Endpoint string `json:"endpoint"`
	} `json:"dhl"`
}

// Parses the config JSON file into a Config struct
func ParseConfigFromJson(fileName string) (conf *Config, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}

	conf = new(Config)
	err = json.NewDecoder(file).Decode(conf)
	
	defer file.Close()
	return
}

// Shortcut for getting the configuration from the default location
func GetConfig() *Config {
	config, err := ParseConfigFromJson("./config.json")
	if err != nil {
		panic(err)
	}
	return config
}
