package controllers

import (
	"encoding/json"
	"mbenz_poc/models"
	"mbenz_poc/utils"

	"github.com/gofiber/fiber/v2"
)

// GetVehicleStatus gets the status of the vehicle's charging level
func GetVehicleStatus(c *fiber.Ctx) error {
	vehicleID := c.Params("vin")
	var vehicleCharge models.VehicleCharge
	if vehicleID != "" {
		req := &utils.Request{
			EndPoint:   "https://restmock.techgig.com/merc/charge_level",
			HTTPMethod: fiber.MethodPost,
			PayLoad:    map[string]string{"vin": vehicleID},
		}
		resp, err := utils.FastHttpManagedBuffers(req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			}) // c.SendStatus(fiber.StatusInternalServerError)
		} else {
			err := json.Unmarshal(resp, &vehicleCharge)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": true,
					"msg":   err.Error(),
				}) // c.SendStatus(fiber.StatusInternalServerError)
			}
		}
	}
	return c.Status(fiber.StatusOK).JSON(vehicleCharge)
}

// GetDistance gets the distance between two POCs
func GetDistance(c *fiber.Ctx) error {
	start := c.Query("source")
	end := c.Query("destination")
	var pocDistance models.DistanceModel
	if start != "" && end != "" {
		req := &utils.Request{
			EndPoint:   "https://restmock.techgig.com/merc/distance",
			HTTPMethod: fiber.MethodPost,
			PayLoad:    map[string]string{"source": start, "destination": end},
		}
		resp, err := utils.FastHttpManagedBuffers(req)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		} else {
			err := json.Unmarshal(resp, &pocDistance)
			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}
	}
	return c.Status(fiber.StatusOK).JSON(pocDistance)
}

func GetStations(c *fiber.Ctx) error {
	start := c.Query("source")
	end := c.Query("destination")
	var chargingStation models.ChargingStations
	if start != "" && end != "" {
		req := &utils.Request{
			EndPoint:   "https://restmock.techgig.com/merc/charging_stations",
			HTTPMethod: fiber.MethodPost,
			PayLoad:    map[string]string{"source": start, "destination": end},
		}
		resp, err := utils.FastHttpManagedBuffers(req)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		} else {
			err := json.Unmarshal(resp, &chargingStation)
			if err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"stations":  chargingStation,
		"nStations": len(chargingStation.Stations),
	})
}
