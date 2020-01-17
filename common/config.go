package common

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	configFile = "./config.yml"

	Global GlobalConfig
	Server ServerConfig
)

// GlobalConfig has all global variables
type GlobalConfig struct {
	Server ServerConfig `yaml:"server"`
}

// ServerConfig contains server related variables
type ServerConfig struct {
	GRPCPort    string  `yaml:"grpcPort"`
	Currency    string  `yaml:"currency"`
	PenPrice    float64 `yanml:"penprice"`
	TshirtPrice float64 `yaml:"tshirtprice"`
	MugPrice    float64 `yaml:"mugprice"`
}

func loadConfig(file string) (GlobalConfig, error) {
	log.Printf("checkout-server: reading config %v\n", file)

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("checkout-server: unexpected reading config failure :%v\n", err)
		return Global, err
	}

	err = yaml.Unmarshal(data, &Global)
	if err != nil {
		log.Printf("checkout-server: unexpected parsing config failure: %v\n", err)
		return Global, err
	}

	Server = Global.Server

	return Global, nil
}
