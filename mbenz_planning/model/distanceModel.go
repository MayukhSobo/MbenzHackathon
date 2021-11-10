package model

import "encoding/json"

type DistanceModel struct {
	Source string `json:"source"`
	Dest string `json:"destination"`
	Distance int `json:"distance"`
}

func (dm *DistanceModel) Unmarshal(rawData []byte) error {
	return json.Unmarshal(rawData, dm)
}