package main

import (
	"fmt"
	"github.com/samuel1992/playlist_migrator/players"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type PlayersConfig struct {
	Spotify players.SpotifyConfig `yaml:"spotify"`
}

func parseConfig() PlayersConfig {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	var config PlayersConfig
	yaml.Unmarshal(data, &config)

	return config
}
