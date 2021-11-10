package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mbenz_planning/model"
	"os"
)

func GetResponse(url string, args []string) ([]byte, []error) {
	agent := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(agent)
	req := agent.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(url)
	for i := 0; i<len(args)-1; i++ {
		req.URI().QueryArgs().Add(args[i], args[i+1])
	}
	if err := agent.Parse(); err != nil {
		return nil, []error{err}
	}

	httpCode, body, err := agent.Bytes()
	if err != nil {
		return nil, err
	}

	if httpCode != fiber.StatusOK {
		return nil, []error{fmt.Errorf("request was not successful. ErrorCode %d", httpCode)}
	}
	return body, nil
}


func GetConcurrentResponse(routeReq *model.RouteRequest) ([][]byte, []error){
	//var results [][]byte
	results := make([][]byte, 2)
	resp1, err := GetResponse(os.Getenv("VEHICLE_INFO_ENDPOINT")+routeReq.VIN, []string{})
	if err != nil {
		return nil, err
	}
	results[0] = resp1
	resp2, err := GetResponse(os.Getenv("POC_DISTANCE_ENDPOINT"), []string{"source", routeReq.Source, "destination", routeReq.Destination})
	if err != nil {
		return nil, err
	}
	results[1] = resp2
	return results, nil
}
