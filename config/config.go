package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	SpaceURL string `json:"space_url"`
	APIKey   string `json:"apikey"`
}

func LoadConfig() (conf Config, err error) {
	raw, err := ioutil.ReadFile("./conf.json")
	if err != nil {
		return conf, err
	}
	json.Unmarshal(raw, &conf)
	return conf, nil
}
