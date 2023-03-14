package players

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

var getPlaylistDetailsResponseJson = `
{
  "collaborative": false,
    "description": "",
    "external_urls": {
      "spotify": "https://open.spotify.com/playlist/3LMA7bVYT6THG5h8Fnc0zL"
    },
    "followers": {
      "href": null,
      "total": 0
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
      "href": "https://api.spotify.com/v1/playlists/3LMA7bVYT6THG5h8Fnc0zL/tracks?offset=0&limit=100",
      "items": [
      {
        "added_at": "2017-09-18T18:05:30Z",
        "added_by": {
          "external_urls": {
            "spotify": "https://open.spotify.com/user/22cb6ttuzh5mjklxeavpv55wy"
          },
          "href": "https://api.spotify.com/v1/users/22cb6ttuzh5mjklxeavpv55wy",
          "id": "22cb6ttuzh5mjklxeavpv55wy",
          "type": "user",
          "uri": "spotify:user:22cb6ttuzh5mjklxeavpv55wy"
        },
        "is_local": false,
        "primary_color": null,
        "track": {
          "album": {
            "album_group": "album",
            "album_type": "album",
            "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/1GhPHrq36VKCY3ucVaZCfo"
              },
              "href": "https://api.spotify.com/v1/artists/1GhPHrq36VKCY3ucVaZCfo",
              "id": "1GhPHrq36VKCY3ucVaZCfo",
              "name": "The Chemical Brothers",
              "type": "artist",
              "uri": "spotify:artist:1GhPHrq36VKCY3ucVaZCfo"
            }
            ],
            "available_markets": [],
            "external_urls": {
              "spotify": "https://open.spotify.com/album/56nVadPbdCs1yGB0AtXSGp"
            },
            "href": "https://api.spotify.com/v1/albums/56nVadPbdCs1yGB0AtXSGp",
            "id": "56nVadPbdCs1yGB0AtXSGp",
            "images": [],
            "is_playable": true,
            "name": "Come With Us",
            "release_date": "2002-01-01",
            "release_date_precision": "day",
            "total_tracks": 10,
            "type": "album",
            "uri": "spotify:album:56nVadPbdCs1yGB0AtXSGp"
          },
          "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/1GhPHrq36VKCY3ucVaZCfo"
            },
            "href": "https://api.spotify.com/v1/artists/1GhPHrq36VKCY3ucVaZCfo",
            "id": "1GhPHrq36VKCY3ucVaZCfo",
            "name": "The Chemical Brothers",
            "type": "artist",
            "uri": "spotify:artist:1GhPHrq36VKCY3ucVaZCfo"
          }
          ],
          "disc_number": 1,
          "duration_ms": 297586,
          "episode": false,
          "explicit": false,
          "external_ids": {
            "isrc": "GBAAA0100912"
          },
          "external_urls": {
            "spotify": "https://open.spotify.com/track/5kTBiVnjq9xKmZL9dNs8zL"
          },
          "href": "https://api.spotify.com/v1/tracks/5kTBiVnjq9xKmZL9dNs8zL",
          "id": "5kTBiVnjq9xKmZL9dNs8zL",
          "is_local": false,
          "name": "Come With Us",
          "popularity": 34,
          "preview_url": null,
          "track": true,
          "track_number": 1,
          "type": "track",
          "uri": "spotify:track:5kTBiVnjq9xKmZL9dNs8zL"
        },
        "video_thumbnail": {
          "url": null
        }
      }
      ],
      "limit": 100,
      "next": null,
      "offset": 0,
      "previous": null,
      "total": 33
    },
    "type": "playlist",
    "uri": "spotify:playlist:3LMA7bVYT6THG5h8Fnc0zL"
}
`

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
