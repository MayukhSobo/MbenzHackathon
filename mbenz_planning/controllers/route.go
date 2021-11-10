package controllers

import (
	"mbenz_planning/model"
)

// PlanRoute plans routes based on the different constraints
func PlanRoute(distance, currentCharge int, stationModel *model.Stations) []model.Station {
	lastDistanceWhenCharged := 0
	var routeStations []model.Station
	for _, station := range stationModel.Stations.AllStations {
		if currentCharge >= distance {
			break
		}
		currentCharge -= station.Dist - lastDistanceWhenCharged
		if currentCharge < 0 {
			// I can't even reach the nearest charging station
			break
		}

		// Stop at current station and recharge
		currentCharge += station.Limit
		routeStations = append(routeStations, station)
		lastDistanceWhenCharged = station.Dist
		distance -= station.Dist
	}
	return routeStations
}