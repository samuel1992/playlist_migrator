package main

import (
	"fmt"
	"github.com/samuel1992/playlist_migrator/players"
	"net/http"
	"os"
)

func main() {
	config := parseConfig()
	player := &players.Spotify{HttpClient: &http.Client{}, Config: config.Spotify}
	player.FetchApiKey()

	playlists, err := player.GetPlaylists()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Playlists:")
	fmt.Println(playlists)
}
