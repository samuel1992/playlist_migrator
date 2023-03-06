package main

import (
	"fmt"
	"github.com/samuel1992/playlist_migrator/players"
	"net/http"
)

func main() {
	config := parseConfig()
	player := players.Spotify{HttpClient: &http.Client{}, Config: config.Spotify}

	fmt.Println("Playlists:")
	fmt.Println(player.GetPlaylists())
}
