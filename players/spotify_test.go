package players

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestGetPlayLists(t *testing.T) {
	getPlaylistsResponseJson, err := ioutil.ReadFile("../testdata/spotify_get_playlists_response.json")
	if err != nil {
		t.Fatal(err)
	}

	httpClient := &MockHttpClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			fakeJson := []byte(getPlaylistsResponseJson)
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader(fakeJson)),
			}, nil
		},
	}

	spotify := &Spotify{
		HttpClient: httpClient,
		Config:     SpotifyConfig{},
	}

	playlists, _ := spotify.GetPlaylists()

	assert.IsType(t, []Playlist{}, playlists)
	assert.Equal(t, 1, len(playlists))
	assert.Equal(t, Playlist{ID: "3LMA7bVYT6THG5h8Fnc0zL", Name: "chemical_brothers", Musics: []Music(nil)}, playlists[0])
}

func TestGetMusics(t *testing.T) {
	getPlaylistDetailsResponseJson, err := ioutil.ReadFile("../testdata/spotify_get_playlist_details_response.json")
	if err != nil {
		t.Fatal(err)
	}

	httpClient := &MockHttpClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			fakeJson := []byte(getPlaylistDetailsResponseJson)
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader(fakeJson)),
			}, nil
		},
	}

	spotify := &Spotify{
		HttpClient: httpClient,
		Config:     SpotifyConfig{},
	}

	musics, _ := spotify.GetMusics("someplaylistid")

	assert.IsType(t, []Music{}, musics)
	assert.Equal(t, 1, len(musics))
	assert.Equal(t, Music{ID: "5kTBiVnjq9xKmZL9dNs8zL", Name: "Come With Us"}, musics[0])
}
