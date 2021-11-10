package models

type DistanceModel struct {
	Source string `json:"source"`
	Destination string `json:"destination"`
	Error string `json:"error,omitempty"`
	Distance int `json:"distance"`
}
