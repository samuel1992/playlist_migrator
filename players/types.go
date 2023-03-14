package players

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Music struct {
	ID   string
	Name string
}

type Playlist struct {
	ID     string
	Name   string
	Musics []Music
}

type Player interface {
	FetchApiKey()
	GetPlaylists() ([]Playlist, error)
	GetMusics(playListId string) ([]Music, error)
}
