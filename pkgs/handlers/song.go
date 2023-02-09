package handlers

import (
	"log"
	models "playlists/pkgs/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetAllSongs(c *fiber.Ctx) error {
	var songs []models.Song
	var err error

	queryError := h.DB.Find(&songs).Error
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
		"data":    songs,
		"message": "Songs fetched successfully",
	})
}

func (h handler) GetSong(c *fiber.Ctx) error {
	var song models.Song
	var id int
	var err error

	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": convErr.Error(),
		})
	}

	queryError := h.DB.Find(&song, id).Error
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
		"data":    song,
		"message": "Song fetched successfully",
	})
}
