package handlers

import (
	"log"
	models "playlists/pkgs/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetAllPlaylists(c *fiber.Ctx) error {
	var playlists []models.Playlist
	var err error

	queryError := h.DB.Find(&playlists).Error
	if queryError != nil {
		err = queryError
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    playlists,
		"message": "Playlists fetched successfully",
	})
}

func (h handler) GetPlaylist(c *fiber.Ctx) error {
	var playlist models.Playlist
	var id int
	var err error

	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": convErr.Error(),
		})
	}

	queryError := h.DB.Find(&playlist, id).Error
	if queryError != nil {
		log.Fatalf("Unable to query database: %v", err)
		err = queryError
	}

	if queryError != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data":    playlist,
		"message": "Playlist fetched successfully",
	})
}
