Migrate playlist from a player to another.

```
Folder structure:
doc/ # Some documentation regarding each player, maybe they have different ways to authenticate, generate api tokens etc
data/
playlist_migrator/
| players/
    | spotify/
        spotify.go
    | deezer/
        deezer.go
    | youtube_music/
        youtube_music.go
    | tidal/
        tidal.go
    music.go
    playlist.go
    main.go
    config.go # maybe read some configuration from a yml file
    configuration.yml #user and password or something, account details (not the safest way to deal with it but)
```

```
# Player class api contract
player = &Spotify{configurations ...}
player.getPlaylist('playlist_id') # returns a list of <Music>
player.createPlaylist(<playlist name>, musicList)
# returns a response with which musics were found in that player and which one was not.
```

```
CLI Interface:

migrator
  --get-playlists
  --get-playlist <playlist id: int>
  --migrate-playlist <playlist_id: int> --from=spotify --to=deezer
```

```
# FANCY IDEAS

# Sso authentication for players:
migrator
  --sso-authenticate --player=spotify
```
