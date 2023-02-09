package handlers

import (
	"log"
	model "playlists/pkgs/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetAllRadioStations(c *fiber.Ctx) error {
	var radioStations []model.RadioStation
	var err error

	queryError := h.DB.Find(&radioStations).Error
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
		"data":    radioStations,
		"message": "Radio stations fetched successfully",
	})
}

func (h handler) GetRadioStation(c *fiber.Ctx) error {
	var radioStation model.RadioStation
	var id int
	var err error

	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": convErr.Error(),
		})
	}

	queryError := h.DB.Find(&radioStation, id).Error
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
		"data":    radioStation,
		"message": "Radio station fetched successfully",
	})
}
