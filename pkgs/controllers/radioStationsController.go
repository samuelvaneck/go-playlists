package controllers

import (
	model "playlists/pkgs/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RadioStationsIndex(c *fiber.Ctx) error {
	radioStations, err := model.RadioStations()
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

func RadioStationById(c *fiber.Ctx) error {
	var id int
	var err error

	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": convErr.Error(),
		})
	}

	radioStation, queryErr := model.RadioStationById(id)
	if queryErr != nil {
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
