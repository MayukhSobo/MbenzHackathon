package model

import "github.com/google/uuid"


type RouteError struct {
	ID int `json:"id"`
	Desc string `json:"description"`
}

type RouteModel struct {
	 TransactionID uuid.UUID `json:"transactionId"`
	 VIN string `json:"vin,omitempty"`
	 Source string `json:"source,omitempty"`
	 Destination string `json:"destination,omitempty"`
	 Distance int `json:"distance,omitempty"`
	 Charge int `json:"currentChargeLevel,omitempty"`
	 ChargeRequired bool `json:"isChargingRequired"`
	 Stations []Station `json:"chargingStations,omitempty"`
	 Errors []RouteError `json:"errors,omitempty"`
}

type ResponseError struct {
	TransactionID uuid.UUID `json:"transactionId"`
	Errors []RouteError `json:"errors"`
}

