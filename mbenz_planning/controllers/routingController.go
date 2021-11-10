package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"mbenz_planning/model"
	"mbenz_planning/utils"
	"os"
)

func sendBadReq(c *fiber.Ctx, err error, fiberStatus int) error {
	var e string
	var res bool
	if err == nil {
		e = "OK"
		res = false
	}else{
		e = err.Error()
		res = true
	}
	return c.Status(fiberStatus).JSON(fiber.Map{
		"error": res,
		"msg": e,
	})
}

func populateResponse(responseModel *model.RouteModel,
	distModel *model.DistanceModel, vehicleInfoModel *model.VehicleInfo, chargeRequired bool, err *model.RouteError) {
	responseModel.TransactionID = uuid.New()
	responseModel.ChargeRequired = chargeRequired
	responseModel.VIN = vehicleInfoModel.VIN
	responseModel.Source = distModel.Source
	responseModel.Destination = distModel.Dest
	responseModel.Distance = distModel.Distance
	responseModel.Charge = vehicleInfoModel.CurrentCharge
	if err != nil {
		responseModel.Errors = append(responseModel.Errors, *err)
	}
}

func GetRoutingDetails(c *fiber.Ctx) error {
	rreq := &model.RouteRequest{}
	if err := c.BodyParser(rreq); err != nil {
		return sendBadReq(c, err, fiber.StatusInternalServerError)
	}

	apiInfo, err := utils.GetConcurrentResponse(rreq)
	if err != nil {
		return sendBadReq(c,
			fmt.Errorf("error while calling the apis"),
			fiber.StatusInternalServerError)
	}

	// Load all the values into structs
	vehicleInfo := &model.VehicleInfo{}
	e := vehicleInfo.Unmarshal(apiInfo[0])
	if e != nil {
		return sendBadReq(c, e, fiber.StatusInternalServerError)
	}
	distanceInfo := &model.DistanceModel{}
	e = distanceInfo.Unmarshal(apiInfo[1])
	if e != nil {
		return sendBadReq(c, e, fiber.StatusInternalServerError)
	}

	responseModel := &model.RouteModel{}


	// ------------------------------------------- //
	// If there is any error in the input
	// Don't call the third API
	if distanceInfo.Distance == 0 || distanceInfo.Source == "" || distanceInfo.Dest == "" {
		var respError []model.RouteError
		respError = append(respError, model.RouteError{ID: 9999, Desc: "Technical Exception"})
		return c.Status(fiber.StatusOK).JSON(model.ResponseError{TransactionID: uuid.New(), Errors: respError})
	}
	// If the current charge is enough to travel the distance
	// Don't call the third API as well
	if distanceInfo.Distance <= vehicleInfo.CurrentCharge {
		populateResponse(responseModel, distanceInfo, vehicleInfo, false, nil)
		return c.Status(fiber.StatusOK).JSON(responseModel)
	}

	// We need to call the third api
	stationResp, err := utils.GetResponse(os.Getenv("POC_STATIONS_ENDPOINT"),
		[]string{"source", rreq.Source, "destination", rreq.Destination})
	if err != nil {
		return sendBadReq(c,
			fmt.Errorf("error calling the station api %v", err),
			fiber.StatusInternalServerError)
	}
	stationModel := &model.Stations{}
	e = stationModel.Unmarshal(stationResp)
	if e != nil {
		return sendBadReq(c, e, fiber.StatusInternalServerError)
	}

	// If the total number of charging station is 0, then we can't
	// reach the destination
	if stationModel.NStations == 0 {
		var respError []model.RouteError
		respError = append(respError, model.RouteError{ID: 8888, Desc: "Unable to reach the destination with the current charge level"})
		return c.Status(fiber.StatusOK).JSON(model.ResponseError{TransactionID: uuid.New(), Errors: respError})
	}

	stations := PlanRoute(distanceInfo.Distance, vehicleInfo.CurrentCharge, stationModel)
	if len(stations) == 0 {
		// We can't reach with the given charge
		populateResponse(responseModel, distanceInfo, vehicleInfo, true,
			&model.RouteError{ID: 8888, Desc: "Unable to reach the destination with the current charge level"})
		return c.Status(fiber.StatusOK).JSON(responseModel)
	}
	populateResponse(responseModel, distanceInfo, vehicleInfo, true, nil)
	responseModel.Stations = stations
	return c.Status(fiber.StatusOK).JSON(responseModel)
}


/*


{
	"vin": "W1K2062161F0033",
	"source": "Home",
	"destination": "Lake"
}
{
	"vin": "W1K2062161F0080",
	"source": "@$%%%",
	"destination": "Airport"
}
{
	"vin": "W1K2062161F0033",
	"source": "Home",
	"destination": "Power Plant"
}

{ "vin": "W1K2062161F0046", "source": "Home", "destination": "Movie Theatre" }
 */