package players

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
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
	HttpClient HTTPClient
	Config     SpotifyConfig
	ApiKey     string
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

type SpotifyPlayListDetailResponse struct {
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
	ExternalUrls  struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  any `json:"href"`
		Total int `json:"total"`
	} `json:"followers"`
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
		Items []struct {
			AddedAt time.Time `json:"added_at"`
			AddedBy struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"added_by"`
			IsLocal      bool `json:"is_local"`
			PrimaryColor any  `json:"primary_color"`
			Track        struct {
				Album struct {
					AlbumType string `json:"album_type"`
					Artists   []struct {
						ExternalUrls struct {
							Spotify string `json:"spotify"`
						} `json:"external_urls"`
						Href string `json:"href"`
						ID   string `json:"id"`
						Name string `json:"name"`
						Type string `json:"type"`
						URI  string `json:"uri"`
					} `json:"artists"`
					AvailableMarkets []string `json:"available_markets"`
					ExternalUrls     struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href   string `json:"href"`
					ID     string `json:"id"`
					Images []struct {
						Height int    `json:"height"`
						URL    string `json:"url"`
						Width  int    `json:"width"`
					} `json:"images"`
					Name                 string `json:"name"`
					ReleaseDate          string `json:"release_date"`
					ReleaseDatePrecision string `json:"release_date_precision"`
					TotalTracks          int    `json:"total_tracks"`
					Type                 string `json:"type"`
					URI                  string `json:"uri"`
				} `json:"album"`
				Artists []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"artists"`
				AvailableMarkets []string `json:"available_markets"`
				DiscNumber       int      `json:"disc_number"`
				DurationMs       int      `json:"duration_ms"`
				Episode          bool     `json:"episode"`
				Explicit         bool     `json:"explicit"`
				ExternalIds      struct {
					Isrc string `json:"isrc"`
				} `json:"external_ids"`
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href        string `json:"href"`
				ID          string `json:"id"`
				IsLocal     bool   `json:"is_local"`
				Name        string `json:"name"`
				Popularity  int    `json:"popularity"`
				PreviewURL  any    `json:"preview_url"`
				Track       bool   `json:"track"`
				TrackNumber int    `json:"track_number"`
				Type        string `json:"type"`
				URI         string `json:"uri"`
			} `json:"track"`
			VideoThumbnail struct {
				URL any `json:"url"`
			} `json:"video_thumbnail"`
		} `json:"items"`
		Limit    int `json:"limit"`
		Next     any `json:"next"`
		Offset   int `json:"offset"`
		Previous any `json:"previous"`
		Total    int `json:"total"`
	} `json:"tracks"`
	Type string `json:"type"`
	URI  string `json:"uri"`
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

	jsonResponseData := &authResponse{}
	err = json.Unmarshal(responseData, jsonResponseData)
	if err != nil {
		fmt.Println(err.Error())
	}

	s.ApiKey = jsonResponseData.AccessToken
}

func (s Spotify) parsePlaylists(data []byte) []Playlist {
	var playlistsResponse SpotifyPlayListsResponse
	json.Unmarshal([]byte(data), &playlistsResponse)

	var playlists []Playlist
	for _, item := range playlistsResponse.Items {
		playlists = append(playlists, Playlist{ID: item.ID, Name: item.Name, Musics: []Music{}})
	}

	return playlists
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

	return s.parsePlaylists(data), nil
}
