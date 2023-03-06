package players

import "net/http"

type Music struct {
	Name string
	ID   string
}

type Playlist struct {
	ID     string
	Name   string
	Musics []Music
}

type Player interface {
	Authenticate(*http.Client, *http.Request)
	GetPlaylists() []Playlist
}
