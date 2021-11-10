package model

import "encoding/json"

type Station struct {
	Name string `json:"name"`
	Dist int `json:"distance"`
	Limit int `json:"limit"`
}
type stations struct {
	Source string `json:"source"`
	Dest string `json:"destination"`
	AllStations []Station `json:"chargingStations"`
}

type Stations struct {
	NStations int `json:"nStations"`
	Stations stations `json:"stations"`
}

func (s *Stations) Unmarshal(rawData []byte) error {
	return json.Unmarshal(rawData, s)
}