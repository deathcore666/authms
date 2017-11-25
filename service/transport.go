package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeLoginEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svcRequest)
		v, err := svc.Login(req.Username, req.Password)
		if err != nil {
			return loginResponse{v, err.Error()}, nil
		}
		return loginResponse{v, ""}, nil
	}
}

func makeRegisterEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svcRequest)
		v, err := svc.Login(req.Username, req.Password)
		if err != nil {
			return loginResponse{v, err.Error()}, nil
		}
		return loginResponse{v, ""}, nil
	}
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request svcRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type svcRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	V   bool   `json:"v"`
	Err string `json:"err,omitempty"`
}

type registerResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}
