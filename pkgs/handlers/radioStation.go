package handlers

import (
	"fmt"
	"log"
	"os/exec"
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
	var out []byte
	out, err := exec.Command("ls", "-l", "../../tmp").Output()
	if err != nil {
		log.Fatal(err)
	}

	// if errors.Is(cmd.Err, exec.ErrDot) {
	// 	cmd.Err = nil
	// }
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }

	// out, err := cmd.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	parsedOut := fmt.Sprintf("%s", out)

	return c.JSON(fiber.Map{
		"success": true,
		"data":    parsedOut,
		"message": "Song recognized successfully",
	})
}
