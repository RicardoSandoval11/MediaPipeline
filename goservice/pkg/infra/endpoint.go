package infra

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type LivenessResponse struct {
	Message string `json:"message"`
}

type ReadinessResponse struct {
	Message string `json:"message"`
}

func MakeLivenessEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := svc.Liveness()
		if err != nil {
			return LivenessResponse{}, err
		}
		return response, nil
	}
}

func MakeReadinessEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response, err := svc.Readiness()
		if err != nil {
			return ReadinessResponse{}, err
		}
		return response, nil
	}
}
