package handlers

import (
	"log"
	models "playlists/pkgs/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetAllArtists(c *fiber.Ctx) error {
	var artists []models.Artist
	var err error

	queryError := h.DB.Find(&artists).Error
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
		"data":    artists,
		"message": "Artists fetched successfully",
	})
}

func (h Handler) GetArtist(c *fiber.Ctx) error {
	var artist models.Artist
	var id int
	var err error

	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": convErr.Error(),
		})
	}

	queryError := h.DB.Find(&artist, id).Error
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
		"data":    artist,
		"message": "Artist fetched successfully",
	})
}
