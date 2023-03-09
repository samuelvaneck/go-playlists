package routes

import (
	"playlists/pkgs/handlers"

	"github.com/gofiber/fiber/v2"
)

func AllRoutes(app *fiber.App, h handlers.Handler) {
	app.Get("/api/v1/radio_stations", h.GetAllRadioStations)
	app.Get("/api/v1/radio_stations/:id", h.GetRadioStation)
	app.Get("/api/v1/radio_stations/:id/recognize", h.RecognizeSong)
	app.Get("/api/v1/playlists", h.GetAllPlaylists)
	app.Get("/api/v1/playlists/:id", h.GetPlaylist)
	app.Get("/api/v1/songs", h.GetAllSongs)
	app.Get("/api/v1/songs/:id", h.GetSong)
	app.Get("/api/v1/artists", h.GetAllArtists)
	app.Get("/api/v1/artists/:id", h.GetArtist)
}
