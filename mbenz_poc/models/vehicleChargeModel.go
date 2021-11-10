package models

type VehicleCharge struct {
	VID string `json:"vin"`
	Charge int `json:"currentChargeLevel"`
	Error string `json:"error,omitempty"`
}
