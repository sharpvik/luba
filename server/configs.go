package main

import (
	"github.com/sharpvik/log-go/v2"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	LubaHost string `yaml:"luba-host"`
	LubaPort int    `yaml:"luba-port"`
}

func (config Config) Log() {
	log.Infof("luba shall run @ %s:%d", config.LubaHost, config.LubaPort)
}

var DefaultConfig = Config{
	LubaHost: "0.0.0.0",
	LubaPort: 7962,
}

func MustConfig(path string) (config Config) {
	file, err := os.Open(path)
	abort(err)
	err = yaml.NewDecoder(file).Decode(&config)
	abort(err)
	return
}

func abort(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
