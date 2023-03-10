package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/samuel1992/playlist_migrator/players"
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

  wg := &sync.WaitGroup{} 
  chMusics := make(chan []players.Music, len(playlists))

  wg.Add(len(playlists))
  for index, playlist := range playlists {
    go getMusics(player, playlist, wg, chMusics)
    playlists[index].Musics = <- chMusics
  }
  wg.Wait()

  for _, playlist := range playlists {
    fmt.Println("new playlist ----->")
    fmt.Printf("%+v\n", playlist)
  }

  fmt.Println("FINISH")
}

 func getMusics(player *players.Spotify, playlist players.Playlist, wg *sync.WaitGroup, chMusics chan []players.Music){
   musics, err := player.GetMusics(playlist.ID)
   if err != nil {
     fmt.Println(err)
   }

   chMusics <- musics
   wg.Done()
 } 
