package model

import "encoding/json"

type VehicleInfo struct {
	VIN string `json:"vin"`
	CurrentCharge int `json:"currentChargeLevel"`
}

func (v *VehicleInfo) Unmarshal(rawData []byte) error {
	return json.Unmarshal(rawData, v)
}