package model

type RouteRequest struct {
	VIN string `json:"vin"`
	Source string `json:"source"`
	Destination string `json:"destination"`
}