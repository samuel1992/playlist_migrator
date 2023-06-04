package main

import (
	"fmt"
	"github.com/samuel1992/playlist_migrator/players"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type PlayersConfig struct {
	Spotify players.SpotifyConfig `yaml:"spotify"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You have to specify an available command:\n - get-playlists")
	}

	command := os.Args[1]

	switch command {
	case "get-playlists":
		getPlaylists()
	}
}

func getPlaylists() {
	config := parseConfig()
	player := &players.Spotify{HttpClient: &http.Client{}, Config: config.Spotify}
	player.Authenticate()

	playlists, err := player.GetPlaylists()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wg := &sync.WaitGroup{}
	for i := range playlists {
		wg.Add(1)
		go fetchMusics(player, &playlists[i], wg)
	}
	wg.Wait()

	for _, playlist := range playlists {
		fmt.Println("PLAYLIST:", playlist.Name)
		fmt.Println("MUSICS:")
		fmt.Printf("%v\n\n", playlist.Musics)
	}
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

func fetchMusics(player players.Player, playlist *players.Playlist, wg *sync.WaitGroup) {
	musics, err := player.GetMusics(playlist.ID)
	if err != nil {
		fmt.Println(err)
	}

	playlist.Musics = musics
	wg.Done()
}
