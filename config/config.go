package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Webserver struct {
		Port string `json:"port"`
		Key string `json:"key"`
	} `json:"webserver"`
	Telegram struct {
		Key string `json:"key"`
		Whitelist []int64 `json:"whitelist"`
	} `json:"telegram"`
	DHL struct {
		Key string `json:"key"`
		Endpoint string `json:"endpoint"`
	} `json:"dhl"`
}

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

func GetConfig() *Config {
	config, err := ParseConfigFromJson("./config.json")
	if err != nil {
		panic(err)
	}
	return config
}
