package models

// ChargingStation represents a single charging station
type ChargingStation struct {
	Name string `json:"name"`
	Distance int `json:"distance"`
	Limit int `json:"limit"`
}

// ChargingStations all the charging stations between two POCs
type ChargingStations struct {
	Source string `json:"source"`
	Destination string `json:"destination"`
	Stations []ChargingStation `json:"chargingStations"`
}
