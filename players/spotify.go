package players

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type SpotifyConfig struct {
	UserID       string `yaml:"user_id"`
	UserApiID    string `yaml:"user_api_id"`
	UserApiToken string `yaml:"user_api_token"`
}

type Spotify struct {
	HttpClient HTTPClient
	Config     SpotifyConfig
	ApiKey     string
}

func (s *Spotify) FetchApiKey() {
	userID := s.Config.UserApiID
	token := s.Config.UserApiToken
	URL := "https://accounts.spotify.com/api/token"

	urlForm := url.Values{}
	urlForm.Set("grant_type", "client_credentials")

	request, err := http.NewRequest("POST", URL, strings.NewReader(urlForm.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(userID, token)
	response, err := s.HttpClient.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	jsonResponseData := &spotifyAuthResponse{}
	err = json.Unmarshal(responseData, jsonResponseData)
	if err != nil {
		fmt.Println(err.Error())
	}

	s.ApiKey = jsonResponseData.AccessToken
}

func parsePlaylists(data []byte) []Playlist {
	var playlistsResponse spotifyPlayListsResponse
	json.Unmarshal([]byte(data), &playlistsResponse)

	var playlists []Playlist
	for _, playlist := range playlistsResponse.Items {
		playlists = append(playlists, Playlist{ID: playlist.ID, Name: playlist.Name})
	}

	return playlists
}

func parsePlaylistMusics(data []byte) []Music {
	var playListDetailResponse spotifyPlayListDetailResponse
	json.Unmarshal([]byte(data), &playListDetailResponse)

	var musics []Music
	for _, music := range playListDetailResponse.Tracks.Items {
		musics = append(musics, Music{ID: music.Track.ID, Name: music.Track.Name})
	}

	return musics
}

func (s Spotify) GetPlaylists() ([]Playlist, error) {
	URL := fmt.Sprintf("https://api.spotify.com/v1/users/%s/playlists", s.Config.UserID)
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", "Bearer "+s.ApiKey)

	response, err := s.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return parsePlaylists(data), nil
}

func (s Spotify) GetMusics(playListId string) ([]Music, error) {
	URL := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s", playListId)

	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", "Bearer "+s.ApiKey)

	response, err := s.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return parsePlaylistMusics(data), nil
}
