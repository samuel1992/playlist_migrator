package main

import (
	"fmt"
	"github.com/samuel1992/playlist_migrator/players"
	"net/http"
	"os"
	"sync"
)

func main() {
	config := parseConfig()
	player := &players.Spotify{HttpClient: &http.Client{}, Config: config.Spotify}
	player.Authenticate()

	playlists, err := player.GetPlaylists()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(playlists))
	for i := range playlists {
		go fetchMusics(player, &playlists[i], wg)
	}
	wg.Wait()

	for _, playlist := range playlists {
		fmt.Println("PLAYLIST:", playlist.Name)
		fmt.Println("MUSICS:")
		fmt.Printf("%v\n\n", playlist.Musics)
	}

	fmt.Println("FINISH")
}

func fetchMusics(player players.Player, playlist *players.Playlist, wg *sync.WaitGroup) {
	musics, err := player.GetMusics(playlist.ID)
	if err != nil {
		fmt.Println(err)
	}

	playlist.Musics = musics
	wg.Done()
}
