package models

type Artist struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name"`
	ID_on_spotify      string `json:"id_on_spotify"`
	Spotify_artist_url string `json:"spotify_url"`
	Spotify_atwork_url string `json:"spotify_atwork_url"`
}
