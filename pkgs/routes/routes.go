package routes

import (
	"playlists/pkgs/handlers"

	"github.com/gofiber/fiber/v2"
)

func AllRoutes(app *fiber.App, h handlers.Handler) {
	app.Get("/api/radio_stations", h.GetAllRadioStations)
	app.Get("/api/radio_stations/:id", h.GetRadioStation)
	app.Get("/api/playlists", h.GetAllPlaylists)
	app.Get("/api/playlists/:id", h.GetPlaylist)
	app.Get("/api/songs", h.GetAllSongs)
	app.Get("/api/songs/:id", h.GetSong)
	app.Get("/api/artists", h.GetAllArtists)
	app.Get("/api/artists/:id", h.GetArtist)
}
