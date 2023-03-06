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

type authResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

type SpotifyConfig struct {
	UserID       string `yaml:"user_id"`
	UserApiID    string `yaml:"user_api_id"`
	UserApiToken string `yaml:"user_api_token"`
}

type Spotify struct {
	HttpClient *http.Client
	Config     SpotifyConfig
}

type SpotifyPlayListsResponse struct {
	Href  string `json:"href"`
	Items []struct {
		Collaborative bool   `json:"collaborative"`
		Description   string `json:"description"`
		ExternalUrls  struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href   string `json:"href"`
		ID     string `json:"id"`
		Images []struct {
			Height int    `json:"height"`
			URL    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name  string `json:"name"`
		Owner struct {
			DisplayName  string `json:"display_name"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"owner"`
		PrimaryColor any    `json:"primary_color"`
		Public       bool   `json:"public"`
		SnapshotID   string `json:"snapshot_id"`
		Tracks       struct {
			Href  string `json:"href"`
			Total int    `json:"total"`
		} `json:"tracks"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"items"`
	Limit    int `json:"limit"`
	Next     any `json:"next"`
	Offset   int `json:"offset"`
	Previous any `json:"previous"`
	Total    int `json:"total"`
}

func (s Spotify) Authenticate(req *http.Request) {
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

	jsonResponseData := &authResponse{}
	err = json.Unmarshal(responseData, jsonResponseData)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Authorization", "Bearer "+jsonResponseData.AccessToken)
}

func (s Spotify) parsePlaylists(data []byte) []Playlist {
	var playlistsResponse SpotifyPlayListsResponse
	json.Unmarshal([]byte(data), &playlistsResponse)

	var playlists []Playlist
	for _, item := range playlistsResponse.Items {
		playlists = append(playlists, Playlist{ID: item.ID, Name: item.Name})
	}

	return playlists
}

func (s Spotify) GetPlaylists() []Playlist {
	URL := fmt.Sprintf("https://api.spotify.com/v1/users/%s/playlists", s.Config.UserID)
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	s.Authenticate(request)

	response, err := s.HttpClient.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return s.parsePlaylists(data)
}
