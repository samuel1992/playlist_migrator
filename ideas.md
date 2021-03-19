Migrate playlist from a player to another.
Folder structure:

playlist_migrator/
-- players/
-- -- spotify/
-- -- -- spotify.py
-- -- deezer/
-- -- -- deezer.py
-- -- youtube_music/
-- -- -- youtube_music.py
-- music.py
-- playlist.py
-- api.py


Which api we expect to a player has?

```
player = Spotify(configurations)
player.get_playlist('playlist_id')
# returns a list of <Music>
player.create_playlist(playlist_name, music_list)
# returns a response with which musics were found in that player and which one was not.
```

Endpoints:

# Players endpoints
GET /players
GET /players/<player_id>
GET /players/<player_id>/playlists

# Playlist in the migrator
GET /playlist/<playlist_id>
POST /playlist/new
BODY:
{ 
  "playlist_name": string,
  "musics": list[music_object]
}

# Insert a new playlist in a player
POST /players/<player_id>/playlists/new
BODY:
{ 
  "playlist_name": string,
  "musics": list[music_object]
}

