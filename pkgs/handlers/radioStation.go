package handlers

import (
	"encoding/json"
	"log"
	models "playlists/pkgs/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetAllRadioStations(c *fiber.Ctx) error {
	var radioStations []models.RadioStation
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

func (h Handler) GetRadioStation(c *fiber.Ctx) error {
	var radioStation models.RadioStation
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

func (h Handler) RecognizeSong(c *fiber.Ctx) error {
	var id int
	var radioStation models.RadioStation
	var err error
	var result map[string]interface{}

	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   convErr.Error(),
			"message": "Unable to convert id to integer",
			"success": false,
		})
	}

	queryError := h.DB.Find(&radioStation, id).Error
	if queryError != nil {
		log.Fatalf("Unable to query database: %v", err)
		err = queryError
	}

	recognizedSong, recErr := radioStation.RecognizeSong()
	if recErr != nil {
		log.Fatalf("Unable to recognize song: %v", recErr)
		err = recErr
	}

	parseErr := json.Unmarshal([]byte(recognizedSong), &result)
	if parseErr != nil {
		log.Fatalf("Unable to parse json: %v", parseErr)
		err = parseErr
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Unable to recognize song",
			"success": false,
		})
	} else {
		return c.JSON(fiber.Map{
			"success": true,
			"data":    result,
			"message": "Song recognized successfully",
		})
	}
}
