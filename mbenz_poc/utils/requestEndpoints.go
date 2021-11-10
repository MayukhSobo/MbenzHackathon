package utils

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type Request struct {
	EndPoint string
	HTTPMethod string
	PayLoad map[string]string
}

func (r *Request) jsonify() ([]byte, error) {
	cj, err := json.Marshal(r.PayLoad)
	if err != nil {
		return []byte{}, nil
	}
	return cj, nil
}

func send(req *fasthttp.Request, resp *fasthttp.Response) ([]byte, error) {
	if err := fasthttp.Do(req, resp); err != nil {
		return []byte{}, fmt.Errorf("client get failed: %s\n", err.Error())
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return []byte{}, fmt.Errorf("expected status code %d but got %d",
			fasthttp.StatusOK, resp.StatusCode())
	}
	if string(resp.Body()) == "" {
		return []byte{}, fmt.Errorf("no response recieved")
	}
	return resp.Body(), nil
}

func get(r *Request) ([]byte, error) {
	// Acquire a request instance
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("GET")
	req.SetRequestURI(r.EndPoint)

	// Acquire a response instance
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	respBody, err := send(req, resp)
	if err != nil {
		return []byte{}, err
	}
	return respBody, nil
}

func post(r *Request) ([]byte, error) {
	// Acquire a request instance
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	jsonPayload, err := r.jsonify()
	if err != nil {
		return []byte{}, fmt.Errorf("unable to marshal the JSON %v", err.Error())
	}
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetBody(jsonPayload)
	req.SetRequestURI(r.EndPoint)

	// Acquire a response instance
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	respBody, err := send(req, resp)
	if err != nil {
		return []byte{}, err
	}
	return respBody, nil
}

// FastHttpManagedBuffers is used to send API requests
func FastHttpManagedBuffers(r *Request) ([]byte, error){
	switch r.HTTPMethod {
	case "GET":
		return get(r)
	case "POST":
		return post(r)
	default:
		return []byte{}, nil
	}
}
