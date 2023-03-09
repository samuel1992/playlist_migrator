package players

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

var getPlaylistsResponseJson = `
{
    "href": "https://api.spotify.com/v1/users/22cb6ttuzh5mjklxeavpv55wy/playlists?offset=0&limit=20",
    "items": [
        {
            "collaborative": false,
            "description": "",
            "external_urls": {
                "spotify": "https://open.spotify.com/playlist/3LMA7bVYT6THG5h8Fnc0zL"
            },
            "href": "https://api.spotify.com/v1/playlists/3LMA7bVYT6THG5h8Fnc0zL",
            "id": "3LMA7bVYT6THG5h8Fnc0zL",
            "images": [
                {
                    "height": 640,
                    "url": "https://i.scdn.co/image/ab67616d0000b273a84e1d8c3a99de8e2e1e6c1a",
                    "width": 640
                }
            ],
            "name": "chemical_brothers",
            "owner": {
                "display_name": "Samuel Dantas",
                "external_urls": {
                    "spotify": "https://open.spotify.com/user/22cb6ttuzh5mjklxeavpv55wy"
                },
                "href": "https://api.spotify.com/v1/users/22cb6ttuzh5mjklxeavpv55wy",
                "id": "22cb6ttuzh5mjklxeavpv55wy",
                "type": "user",
                "uri": "spotify:user:22cb6ttuzh5mjklxeavpv55wy"
            },
            "primary_color": null,
            "public": true,
            "snapshot_id": "NSw0OTdlYjFlNTE2MmNjY2M5ODkwYzViZTk5YWVjYTYxYTIzN2NkZDk4",
            "tracks": {
                "href": "https://api.spotify.com/v1/playlists/3LMA7bVYT6THG5h8Fnc0zL/tracks",
                "total": 33
            },
            "type": "playlist",
            "uri": "spotify:playlist:3LMA7bVYT6THG5h8Fnc0zL"
        }
    ],
    "limit": 20,
    "next": null,
    "offset": 0,
    "previous": null,
    "total": 6
}`

type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestGetPlayLists(t *testing.T) {
	httpClient := &MockHttpClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			fakeJson := []byte(getPlaylistsResponseJson)
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader(fakeJson)),
			}, nil
		},
	}

	spotify := Spotify{
		HttpClient: httpClient,
		Config:     SpotifyConfig{},
	}

	playlists, _ := spotify.GetPlaylists()

	assert.IsType(t, []Playlist{}, playlists)
	assert.Equal(t, 1, len(playlists))
	assert.Equal(t, Playlist{ID: "3LMA7bVYT6THG5h8Fnc0zL", Name: "chemical_brothers", Musics: []Music{}}, playlists[0])
}
